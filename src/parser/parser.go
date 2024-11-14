package parser

import (
	"af/src/ast"
	"af/src/lexer"
	"af/src/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	parser := &Parser{l: l}

	// calling twice to set curToken and peekToken
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	// TODO: start parsing statements wit parser.nextToken() and append to program.Statements
	return program
}
