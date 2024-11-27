package ast

import (
	"af/src/token"
	"bytes"
)

type Node interface {
	TokenLiteral() string
	PrintAsString() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) PrintAsString() string {
	var out bytes.Buffer

	for _, stmt := range p.Statements {
		out.WriteString(stmt.PrintAsString())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) PrintAsString() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.PrintAsString() + " ")
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.PrintAsString())
	}
	out.WriteString(";")
	return out.String()
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) PrintAsString() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) PrintAsString() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.PrintAsString())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) PrintAsString() string {
	if es.Expression != nil {
		return es.Expression.PrintAsString()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) PrintAsString() string {
	return il.Token.Literal
}
