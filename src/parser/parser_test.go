package parser

import (
	"af/src/ast"
	"af/src/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10.5;
	let foo = 15;
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nill")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not have 3 statementes, got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{expectedIdentifier: "x"},
		{expectedIdentifier: "y"},
		{expectedIdentifier: "foo"},
	}
	for i, test := range tests {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, test.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.GetErrors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, errorMsg := range errors {
		t.Errorf("Parser error: %q", errorMsg)
	}
	t.FailNow()
}
func testLetStatements(t *testing.T, s ast.Statement, expected string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Expected 'let', got %q", s.TokenLiteral())
		return false
	}
	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStatement.Name.Value != expected {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", expected, letStatement.Name.Value)
		return false
	}
	if letStatement.Name.TokenLiteral() != expected {
		t.Errorf("letStatement.Name.TokenLiteral() not '%s'. got=%s", expected, letStatement.Name.TokenLiteral())
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10.5;
	return "hello";
	`
	lexer := lexer.NewLexer(input)
	parser := NewParser(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatalf("ParseProgram() returned nill")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not have 3 statementes, got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestParseIdentifiersExpressions(t *testing.T) {
	input := `foobar;
	hello;
	world;`
	lexer := lexer.NewLexer(input)
	parser := NewParser(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	tests := []struct {
		identifierName string
	}{
		{identifierName: "foobar"},
		{identifierName: "hello"},
		{identifierName: "world"},
	}
	if program == nil {
		t.Fatalf("ParseProgram() returned nill")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not have 3 statementes, got %d", len(program.Statements))
	}

	for i, test := range tests {
		identStmt, ok := program.Statements[i].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("stmt not *ast.ExpressionStatement. got=%T", identStmt)
			continue
		}
		if identStmt.TokenLiteral() != test.identifierName {
			t.Errorf("returnStmt.TokenLiteral not '%q', got %q",
				test.identifierName, identStmt.TokenLiteral())
		}
	}
}

func TestParseIntExpressions(t *testing.T) {
	input := `5;
	10;
	15;`
	lexer := lexer.NewLexer(input)
	parser := NewParser(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	tests := []struct {
		integerValue   int64
		integerLiteral string
	}{
		{integerValue: 5, integerLiteral: "5"},
		{integerValue: 10, integerLiteral: "10"},
		{integerValue: 15, integerLiteral: "15"},
	}
	if program == nil {
		t.Fatalf("ParseProgram() returned nill")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not have 3 statementes, got %d", len(program.Statements))
	}

	for i, test := range tests {
		identStmt, ok := program.Statements[i].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("stmt not *ast.ExpressionStatement. got=%T", identStmt)
			continue
		}
		literal, ok := identStmt.Expression.(*ast.IntegerLiteral)

		if !ok {
			t.Fatalf("exp not *ast.IntegerLiteral. got=%T", identStmt.Expression)
		}

		if literal.TokenLiteral() != test.integerLiteral {
			t.Errorf("literal.TokenLiteral not %s. got=%s", test.integerLiteral,
				literal.TokenLiteral())
		}

		if literal.Value != test.integerValue {
			t.Errorf("literal.Value not %d. got=%d", test.integerValue,
				literal.Value)
		}
	}
}

// TODO: print program statements as text
// requires: parsing expression
func TestPrintAsStringStatements(t *testing.T) {
	input := `
	return 5;
	return 10.5;
	return "hello";
	`
	lexer := lexer.NewLexer(input)
	parser := NewParser(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatalf("ParseProgram() returned nill")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not have 3 statementes, got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}
