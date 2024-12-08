package parser

import (
	"af/src/ast"
	"af/src/lexer"
	"af/src/token"
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUAL       // ==
	LESSGREATER // < > <= >=
	SUM         // + -
	PRODUCT     // * /
	PREFIX      // !true -5
	CALL        // add()
)

var precedences = map[token.TokenType]int{
	token.EQUAL:     EQUAL,
	token.NOT_EQUAL: EQUAL,
	token.LT:        LESSGREATER,
	token.GT:        LESSGREATER,
	token.PLUS:      SUM,
	token.MINUS:     SUM,
	token.SLASH:     PRODUCT,
	token.ASTERISK:  PRODUCT,
	token.MOD:       PRODUCT,
}

type (
	prefixParseFN = func() ast.Expression
	infixParseFN  = func(ast.Expression) ast.Expression
)

type Parser struct {
	l               *lexer.Lexer
	curToken        token.Token
	peekToken       token.Token
	errors          []string
	prefixParserFns map[token.TokenType]prefixParseFN
	infixParserFns  map[token.TokenType]infixParseFN
}

func NewParser(l *lexer.Lexer) *Parser {
	parser := &Parser{
		l:      l,
		errors: []string{},
	}

	parser.prefixParserFns = make(map[token.TokenType]prefixParseFN)
	parser.infixParserFns = make(map[token.TokenType]infixParseFN)
	parser.prefixParserFns[token.IDENT] = parser.parseIdentifier
	parser.prefixParserFns[token.INT] = parser.parseInt
	parser.prefixParserFns[token.FLOAT] = parser.parseFloat
	parser.prefixParserFns[token.BANG] = parser.parsePrefixExpression
	parser.prefixParserFns[token.MINUS] = parser.parsePrefixExpression
	parser.prefixParserFns[token.TRUE] = parser.parseBoolean
	parser.prefixParserFns[token.FALSE] = parser.parseBoolean
	parser.prefixParserFns[token.LPAREN] = parser.parseGroupedExpression

	parser.infixParserFns[token.PLUS] = parser.parseInfixExpression
	parser.infixParserFns[token.MINUS] = parser.parseInfixExpression
	parser.infixParserFns[token.EQUAL] = parser.parseInfixExpression
	parser.infixParserFns[token.NOT_EQUAL] = parser.parseInfixExpression
	parser.infixParserFns[token.ASTERISK] = parser.parseInfixExpression
	parser.infixParserFns[token.SLASH] = parser.parseInfixExpression
	parser.infixParserFns[token.LT] = parser.parseInfixExpression
	parser.infixParserFns[token.GT] = parser.parseInfixExpression
	parser.infixParserFns[token.MOD] = parser.parseInfixExpression

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
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	// TODO: skipping the expressions until a semicolon is found
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST) // precedence will be used to evaluate correctly expressions
	// useful for REPL to evaluate  for example 5+5 without semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParserFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExpression := prefix()
	for p.peekToken.Type != token.SEMICOLON && precedence < p.peekTokenPrecedence() {
		infix := p.infixParserFns[p.peekToken.Type]
		if infix == nil {
			return leftExpression
		}
		p.nextToken()
		leftExpression = infix(leftExpression)
	}
	return leftExpression
}
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseInt() ast.Expression {
	il := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		errorMsg := fmt.Sprintf("Could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, errorMsg)
		return nil
	}

	il.Value = value
	return il
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}

func (p *Parser) parseFloat() ast.Expression {
	il := &ast.FloatLiteral{Token: p.curToken}
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		errorMsg := fmt.Sprintf("Could not parse %q as float", p.curToken.Literal)
		p.errors = append(p.errors, errorMsg)
		return nil
	}

	il.Value = value
	return il
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	pe := &ast.PrefixExpression{Token: p.curToken, Operator: p.curToken.Literal}

	p.nextToken()
	e := p.parseExpression(PREFIX)
	pe.Right = e
	return pe
}

func (p *Parser) parseInfixExpression(e ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{Token: p.curToken, Operator: p.curToken.Literal, Left: e}
	currentPrecedence := p.curTokenPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(currentPrecedence)
	return expression
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noInfixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no infix parse function for %s found", t)
	p.errors = append(p.errors, msg)
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

func (p *Parser) curTokenPrecedence() int {
	if precedence, ok := precedences[p.curToken.Type]; ok {
		return precedence
	}
	return LOWEST
}

func (p *Parser) peekTokenPrecedence() int {
	if precedence, ok := precedences[p.peekToken.Type]; ok {
		return precedence
	}
	return LOWEST
}
