package parser

import (
	"testing"

	"go-interpreter/ast"
	"go-interpreter/lexer"
)

func Test_LetStatements(t *testing.T) {
	t.Run("should parse the let statement successfully", func(t *testing.T) {
		input := `
let x = 5;
let y = 10;
let foobar = 121212;
`
		l := lexer.New(input)
		p := New(l)

		program := p.ParseProgram()
		checkParseErrors(t, p)
		if program == nil {
			t.Fatalf("ParseProgram() returned nil")
		}
		if len(program.Statements) != 3 {
			t.Fatalf("program.Statements doesn't contain 3 statements")
		}

		tests := []struct {
			expressionIdentifier string
		}{
			{"x"},
			{"y"},
			{"foobar"},
		}

		for i, tt := range tests {
			stmt := program.Statements[i]
			if !testLetStatement(t, stmt, tt.expressionIdentifier) {
				return
			}
		}
	})

	t.Run("should fail to parse the let statement", func(t *testing.T) {
		input := `
let x 5;
let = 10;
let 123;`

		l := lexer.New(input)
		p := New(l)

		program := p.ParseProgram()
		checkParseErrors(t, p)
		if program == nil {
			t.Fatalf("ParseProgram() returned nil")
		}
		if len(program.Statements) != 3 {
			t.Fatalf("program.Statements doesn't contain 3 statements")
		}

		tests := []struct {
			expressionIdentifier string
		}{
			{"x"},
			{"y"},
			{"foobar"},
		}

		for i, tt := range tests {
			stmt := program.Statements[i]
			if !testLetStatement(t, stmt, tt.expressionIdentifier) {
				return
			}
		}
	})

}

func Test_ReturnStatement(t *testing.T) {
	t.Run("should parse return statements successfully", func(t *testing.T) {
		input := `
return 1;
return 2;
return 3;`

		l := lexer.New(input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)
		if program == nil {
			t.Fatalf("Program have no statements")
		}
		if len(program.Statements) != 3 {
			t.Fatalf("Program doesnt contain 3 statements")
		}

		for _, val := range program.Statements {
			stmt, ok := val.(*ast.ReturnStatement)
			if !ok {
				t.Fatalf("Return statement assert failed, got type %T", val)
			}
			if stmt.TokenLiteral() != "return" {
				t.Fatalf("return statement token is not 'return', but got value of %s", stmt.TokenLiteral())
			}
		}

	})

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral is not let, got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("assertion to LetStatement failed, got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
