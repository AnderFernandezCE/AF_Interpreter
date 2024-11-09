package lexer

import (
	"af/src/token"
	"testing"
)

func TestTokens(t *testing.T) {
	input := `let numVariable = 5.6;
	let otherVariable = 10;

	fn addFn (x, y){
		x + y;
	}
	let result = addFn(numVariable, otherVariable);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "numVariable"},
		{token.ASSIGN, "="},
		{token.FLOAT, "5.6"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "otherVariable"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "addFn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "addFn"},
		{token.LPAREN, "("},
		{token.IDENT, "numVariable"},
		{token.COMMA, ","},
		{token.IDENT, "otherVariable"},
		{token.RPAREN, ")"},
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
