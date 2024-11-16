package parser

import (
	"af/src/ast"
	"af/src/lexer"
	"af/src/token"
	"fmt"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func NewParser(l *lexer.Lexer) *Parser {
	parser := &Parser{
		l:      l,
		errors: []string{},
	}

	// calling twice to set curToken and peekToken
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (p *Parser) GetErrors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	emsg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, emsg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	// TODO: start parsing statements wit parser.nextToken() and append to program.Statements
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	identifier := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	stmt.Name = identifier

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping the expressions until a semicolon is found
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	// TODO: skipping the expressions until a semicolon is found
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// checks next token and advances one token, will be useful for handling errors
func (p *Parser) expectPeek(token token.TokenType) bool {
	if p.peekTokenIs(token) {
		p.nextToken()
		return true
	}
	p.peekError(token)
	return false
}

func (p *Parser) peekTokenIs(token token.TokenType) bool {
	return p.peekToken.Type == token
}
func (p *Parser) curTokenIs(token token.TokenType) bool {
	return p.curToken.Type == token
}
