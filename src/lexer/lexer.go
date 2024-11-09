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
	l.skipWhitespaces()

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
		if isLetter(l.currentValue) {
			tok.Literal = l.readIndentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isNumber(l.currentValue) {
			tok.Literal = l.readNumber()
			tok.Type = token.LookUpNumberType(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentValue)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (l *Lexer) readIndentifier() string {
	var initialPosition = l.position
	for isLetter(l.currentValue) {
		l.readChar()
	}
	return l.input[initialPosition:l.position]
}

func (l *Lexer) readNumber() string {
	var initialPosition = l.position
	var isFloatNumber = false
	for isNumber(l.currentValue) {
		l.readChar()
	}
	if l.currentValue == '.' {
		isFloatNumber = true
		l.readChar()
	}
	if isFloatNumber {
		for isNumber(l.currentValue) {
			l.readChar()
		}
	}
	return l.input[initialPosition:l.position]
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) skipWhitespaces() {
	for l.currentValue == ' ' || l.currentValue == '\n' || l.currentValue == '\t' || l.currentValue == '\r' {
		l.readChar()
	}
}
