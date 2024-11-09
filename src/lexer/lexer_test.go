package lexer

import (
	"af/src/token"
	"testing"
)

func TestTokens(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	lexer := NewLexer(input)

	for index, testValue := range tests {
		tok := lexer.NextToken()

		if tok.Type != testValue.expectedType {
			t.Fatalf("Tests [%d] - tokentype wrong. Expected=%v , got=%v", index, testValue.expectedType, tok.Type)
		}

		if tok.Literal != testValue.expectedLiteral {
			t.Fatalf("Tests [%d] - literal wrong. Expected=%v , got=%v", index, testValue.expectedLiteral, tok.Literal)
		}
	}

}
