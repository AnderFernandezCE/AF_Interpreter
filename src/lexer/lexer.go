package lexer

import "af/src/token"

type Lexer struct {
	input        string
	position     int  // position in the input
	nextPosition int  // next position
	currentValue byte // char at the current position
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // sets initial values for position, nextPosition and currentValue
	return l
}

// function to read current char and advance position
func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.currentValue = 0
	} else {
		l.currentValue = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.currentValue {
	case '+':
		tok = newToken(token.PLUS, l.currentValue)
	case '(':
		tok = newToken(token.LPAREN, l.currentValue)
	case ')':
		tok = newToken(token.RPAREN, l.currentValue)
	case '{':
		tok = newToken(token.LBRACE, l.currentValue)
	case '}':
		tok = newToken(token.RBRACE, l.currentValue)
	case ',':
		tok = newToken(token.COMMA, l.currentValue)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentValue)
	case '=':
		tok = newToken(token.ASSIGN, l.currentValue)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		tok.Type = token.ILLEGAL
		tok.Literal = ""
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
