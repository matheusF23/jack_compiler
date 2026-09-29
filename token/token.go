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
}

func NewToken(tokenType TokenType, lexeme string) Token {
	return Token{
		Type:   tokenType,
		Lexeme: lexeme,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("<%s>%s</%s>", t.Type, t.Lexeme, t.Type)
}
