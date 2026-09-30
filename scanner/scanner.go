package scanner

import (
	"fmt"
	"unicode"

	"jack_compiler/token"
)

type Scanner struct {
	input   []byte
	current int
	line    int
}

var keywords = map[string]token.TokenType{
	"class":       token.CLASS,
	"constructor": token.CONSTRUCTOR,
	"function":    token.FUNCTION,
	"method":      token.METHOD,
	"field":       token.FIELD,
	"static":      token.STATIC,
	"var":         token.VAR,
	"int":         token.INT,
	"char":        token.CHAR,
	"boolean":     token.BOOLEAN,
	"void":        token.VOID,
	"true":        token.TRUE,
	"false":       token.FALSE,
	"null":        token.NULL,
	"this":        token.THIS,
	"let":         token.LET,
	"do":          token.DO,
	"if":          token.IF,
	"else":        token.ELSE,
	"while":       token.WHILE,
	"return":      token.RETURN,
}

func NewScanner(input []byte) *Scanner {
	return &Scanner{
		input: input,
		line:  1,
	}
}

func (s *Scanner) peek() byte {
	if s.current < len(s.input) {
		return s.input[s.current]
	}

	return '\x00'
}

func (s *Scanner) peekNext() byte {
	next := s.current + 1

	if next < len(s.input) {
		return s.input[next]
	}

	return '\x00'
}

func (s *Scanner) advance() {
	ch := s.peek()

	if ch != '\x00' {
		s.current++
	}
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || unicode.IsDigit(rune(c))
}

func (s *Scanner) number() token.Token {
	start := s.current

	for unicode.IsDigit(rune(s.peek())) {
		s.advance()
	}

	n := string(s.input[start:s.current])

	return token.NewToken(token.NUMBER, n, s.line)
}

func (s *Scanner) identifier() token.Token {
	start := s.current

	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	id := string(s.input[start:s.current])

	typeToken, ok := keywords[id]

	if !ok {
		typeToken = token.IDENT
	}

	return token.NewToken(typeToken, id, s.line)
}

func (s *Scanner) string() token.Token {
	startLine := s.line
	s.advance() // Skip the opening quote
	start := s.current
	for s.peek() != '"' && s.peek() != '\x00' {
		s.advance()
	}
	str := string(s.input[start:s.current])

	s.advance() // consume the closing quote

	return token.NewToken(token.STRING, str, startLine)
}

func (s *Scanner) NextToken() token.Token {
	s.skipWhitespace()

	ch := s.peek()

	if isAlpha(ch) {
		return s.identifier()
	}

	if ch == '0' {
		s.advance()

		return token.NewToken(
			token.NUMBER,
			string(ch),
			s.line,
		)
	} else if unicode.IsDigit(rune(ch)) {
		return s.number()
	}

	switch ch {
	case '+':
		s.advance()
		return token.NewToken(token.PLUS, "+", s.line)

	case '-':
		s.advance()
		return token.NewToken(token.MINUS, "-", s.line)

	case '*':
		s.advance()
		return token.NewToken(token.ASTERISK, "*", s.line)

	case '/':
		if s.peekNext() == '/' {
			s.skipLineComments()
			return s.NextToken()
		} else if s.peekNext() == '*' {
			s.skipBlockComments()
			return s.NextToken()
		} else {
			s.advance()
			return token.NewToken(token.SLASH, "/", s.line)
		}

	case '=':
		s.advance()
		return token.NewToken(token.EQ, "=", s.line)

	case ';':
		s.advance()
		return token.NewToken(token.SEMICOLON, ";", s.line)

	case '"':
		return s.string()

	case '\x00':
		return token.NewToken(token.EOF, "EOF", s.line)

	default:
		panic(fmt.Sprintf("lexical error at %c", ch))
	}
}

func (s *Scanner) skipWhitespace() {
	ch := s.peek()

	for ch == ' ' || ch == '\r' || ch == '\t' || ch == '\n' {
		if ch == '\n' {
			s.line++
		}

		s.advance()
		ch = s.peek()
	}
}

func (s *Scanner) skipLineComments() {
	for s.peek() != '\n' && s.peek() != '\x00' {
		s.advance()
	}
}

func (s *Scanner) skipBlockComments() {
	s.advance()
	s.advance()
	for {
		if s.peek() == '\x00' {
			panic("Comentário de bloco não fechado")
		}
		if s.peek() == '*' && s.peekNext() == '/' {
			s.advance()
			s.advance()
			return
		}
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
}
