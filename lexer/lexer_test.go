package lexer

import (
	"testing"

	"github.com/KishorPokharel/calculator/token"
)

func TestNextToken(t *testing.T) {
	input := `()+-*/`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.PLUS, "+"},
		{token.SUBTRACT, "-"},
		{token.MULTIPLY, "*"},
		{token.DIVIDE, "/"},
	}

	l := New(input)
	for i, tt := range tests {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong expected = %q got = %q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong expected = %q got = %q", i, tt.expectedLiteral, token.Literal)
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `55 + 40`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NUMBER, "55"},
		{token.PLUS, "+"},
		{token.NUMBER, "40"},
	}

	l := New(input)
	for i, tt := range tests {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("2:tests[%d]-tokentype wrong! expected %q got %q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf("2:tests[%d] - literal wrong expected = %q got = %q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
