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

	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
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
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQUAL, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQUAL, "!="},
		{token.INT, "9"},
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
