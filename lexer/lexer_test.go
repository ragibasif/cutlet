package lexer

import (
	"testing"

	"github.com/ragibasif/cutlet/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fun(x, y) {
x + y;
};
let result = add(five, ten);
`
	tests := []struct {
		expectedType    token.TokenKind
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGNMENT, "="},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGNMENT, "="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGNMENT, "="},
		{token.FUN, "fun"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGNMENT, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Kind != tt.expectedType {
			t.Fatalf("tests[%d] - tokenkind wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Kind)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
