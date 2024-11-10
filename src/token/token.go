package token

import "strings"

type TokenType string

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	MOD      = "%"
	BANG     = "!"
	ASTERISK = "*"

	LT        = "<"
	GT        = ">"
	EQUAL     = "=="
	NOT_EQUAL = "!="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}
