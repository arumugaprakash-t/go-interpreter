package lexer

import (
	"go-interpreter/token"
	"testing"
)

func Test_NextToken(t *testing.T) {

	t.Run("should run successfully with given input", func(t *testing.T) {
		input := "=,:(){}"

		tests := []struct {
			expectedType   token.TokenType
			expectedString string
		}{
			{token.EQUAL, "="},
			{token.COMMA, ","},
			{token.COLON, ":"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
					i, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedString {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i, tt.expectedString, tok.Literal)
			}
		}

	})

	t.Run("should run successfully for given function input", func(t *testing.T) {
		input := `let five = 5;
			let ten = 10;
   			let add = fn(x, y) {
    			 x + y;
			};
   			let result = add(five, ten);`

		tests := []struct {
			expectedType   token.TokenType
			expectedString string
		}{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.EQUAL, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.EQUAL, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.EQUAL, "="},
			{token.FUNCTION, "fn"},
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
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.EQUAL, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
					i, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedString {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i, tt.expectedString, tok.Literal)
			}
		}
	})
}
