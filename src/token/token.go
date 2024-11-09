package token

import "strings"

type TokenType string

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}

func LookUpNumberType(ident string) TokenType {
	if strings.Contains(ident, ".") {
		return FLOAT
	}
	return INT
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Variables
	IDENT = "IDENT"
	INT   = "INT"
	FLOAT = "FLOAT"

	// Operators
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	DVISION = "/"
	MOD     = "%"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}
