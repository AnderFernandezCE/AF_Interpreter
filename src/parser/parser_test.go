package parser

import (
	"af/src/ast"
	"af/src/lexer"
	"fmt"
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

func TestParseIBooleanExpressions(t *testing.T) {
	input := `true;
	false;`
	lexer := lexer.NewLexer(input)
	parser := NewParser(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	tests := []struct {
		value string
	}{
		{value: "true"},
		{value: "false"},
	}
	if program == nil {
		t.Fatalf("ParseProgram() returned nill")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not have 2 statementes, got %d", len(program.Statements))
	}

	for i, test := range tests {
		boolStmt, ok := program.Statements[i].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("stmt not *ast.ExpressionStatement. got=%T", boolStmt)
			continue
		}
		if boolStmt.TokenLiteral() != test.value {
			t.Errorf("boolStmt.TokenLiteral not '%q', got %q",
				test.value, boolStmt.TokenLiteral())
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

func TestParseFloatExpressions(t *testing.T) {
	input := `5.0;
	10.5;
	15.0;`
	lexer := lexer.NewLexer(input)
	parser := NewParser(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	tests := []struct {
		value          float64
		integerLiteral string
	}{
		{value: 5, integerLiteral: "5.0"},
		{value: 10.5, integerLiteral: "10.5"},
		{value: 15, integerLiteral: "15.0"},
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
		literal, ok := identStmt.Expression.(*ast.FloatLiteral)

		if !ok {
			t.Fatalf("exp not *ast.FloatLiteral. got=%T", identStmt.Expression)
		}

		if literal.TokenLiteral() != test.integerLiteral {
			t.Errorf("literal.TokenLiteral not %s. got=%s", test.integerLiteral,
				literal.TokenLiteral())
		}

		if literal.Value != test.value {
			t.Errorf("literal.Value not %f. got=%f", test.value,
				literal.Value)
		}
	}
}

func TestParsePrefixExpressions(t *testing.T) {
	input := `!5.0;
	-10.78;
	-15.0;
	!true;
	!false;`
	lexer := lexer.NewLexer(input)
	parser := NewParser(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	tests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{input: "!5.0", value: 5.0, operator: "!"},
		{input: "-10.78", value: 10.78, operator: "-"},
		{input: "-15.0", value: 15.0, operator: "-"},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}
	if program == nil {
		t.Fatalf("ParseProgram() returned nill")
	}

	if len(program.Statements) != 5 {
		t.Fatalf("program.Statements does not have 5 statementes, got %d", len(program.Statements))
	}

	for i, test := range tests {
		expression, ok := program.Statements[i].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("stmt not *ast.ExpressionStatement. got=%T", expression)
			continue
		}
		prefix, ok := expression.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf("exp not *ast.PrefixExpression. got=%T", prefix)
		}

		if prefix.Operator != test.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				test.operator, prefix.Operator)
		}

		if !testLiteralExpression(t, prefix.Right, test.value) {
			return
		}

	}
}

func TestParseInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}
	for _, tt := range infixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		if !testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}
	return true
}

func testFloatLiteral(t *testing.T, il ast.Expression, value float64) bool {
	float, ok := il.(*ast.FloatLiteral)
	if !ok {
		t.Errorf("fl not *ast.FloatLiteral. got=%T", il)
		return false
	}
	if float.Value != value {
		t.Errorf("float.Value not %f. got=%f", value, float.Value)
		return false
	}
	return true
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

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case float64:
		return testFloatLiteral(t, exp, float64(v))
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}
	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}
	if bo.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bo.TokenLiteral not %t. got=%s",
			value, bo.TokenLiteral())
		return false
	}
	return true
}
