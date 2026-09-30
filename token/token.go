package token

import "fmt"

type TokenType string

const (
	// Symbols
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	LBRACE    TokenType = "LBRACE"
	RBRACE    TokenType = "RBRACE"
	LBRACKET  TokenType = "LBRACKET"
	RBRACKET  TokenType = "RBRACKET"
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"
	DOT       TokenType = "DOT"
	PLUS      TokenType = "PLUS"
	MINUS     TokenType = "MINUS"
	ASTERISK  TokenType = "ASTERISK"
	SLASH     TokenType = "SLASH"
	AND       TokenType = "AND"
	OR        TokenType = "OR"
	NOT       TokenType = "NOT"
	LT        TokenType = "LT"
	GT        TokenType = "GT"
	EQ        TokenType = "EQ"

	// Literals
	NUMBER TokenType = "NUMBER"
	IDENT  TokenType = "IDENT"
	STRING TokenType = "STRING"

	// Keywords
	CLASS       TokenType = "CLASS"
	CONSTRUCTOR TokenType = "CONSTRUCTOR"
	FUNCTION    TokenType = "FUNCTION"
	METHOD      TokenType = "METHOD"
	FIELD       TokenType = "FIELD"
	STATIC      TokenType = "STATIC"
	VAR         TokenType = "VAR"
	INT         TokenType = "INT"
	CHAR        TokenType = "CHAR"
	BOOLEAN     TokenType = "BOOLEAN"
	VOID        TokenType = "VOID"
	TRUE        TokenType = "TRUE"
	FALSE       TokenType = "FALSE"
	NULL        TokenType = "NULL"
	THIS        TokenType = "THIS"
	LET         TokenType = "LET"
	DO          TokenType = "DO"
	IF          TokenType = "IF"
	ELSE        TokenType = "ELSE"
	WHILE       TokenType = "WHILE"
	RETURN      TokenType = "RETURN"

	EOF TokenType = "EOF"
)

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
}

func NewToken(tokenType TokenType, lexeme string, line int) Token {
	return Token{
		Type:   tokenType,
		Lexeme: lexeme,
		Line:   line,
	}
}

func isSymbol(t TokenType) bool {
	switch t {
	case LPAREN, RPAREN,
		LBRACE, RBRACE,
		LBRACKET, RBRACKET,
		COMMA, SEMICOLON,
		DOT, PLUS, MINUS,
		ASTERISK, SLASH,
		AND, OR, NOT,
		LT, GT, EQ:
		return true
	default:
		return false
	}
}

func (t Token) String() string {
	var category string

	switch t.Type {
	case NUMBER:
		category = "integerConstant"
	case IDENT:
		category = "identifier"
	case STRING:
		category = "stringConstant"
	default:
		if isSymbol(t.Type) {
			category = "symbol"
		} else {
			category = "keyword"
		}
	}

	value := t.Lexeme

	if category == "symbol" {
		switch value {
		case ">":
			value = "&gt;"
		case "<":
			value = "&lt;"
		case `"`:
			value = "&quot;"
		case "&":
			value = "&amp;"
		}
	}

	return fmt.Sprintf("<%s> %s </%s>", category, value, category)
}
