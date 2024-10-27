package parser

import (
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input          string
		expectedTokens []struct {
			token   Token
			literal string
		}
	}{
		{
			input: "func myFunction(a, b) { return a + b; }",
			expectedTokens: []struct {
				token   Token
				literal string
			}{
				{token: FUNC, literal: "func"},
				{token: IDENT, literal: "myFunction"},
				{token: LPAREN, literal: "("},
				{token: IDENT, literal: "a"},
				{token: COMMA, literal: ","},
				{token: IDENT, literal: "b"},
				{token: RPAREN, literal: ")"},
				{token: LBRACE, literal: "{"},
				{token: IDENT, literal: "return"},
				{token: IDENT, literal: "a"},
				{token: ADD, literal: "+"},
				{token: IDENT, literal: "b"},
				{token: SEMI, literal: ";"},
				{token: RBRACE, literal: "}"},
				{token: EOF, literal: ""},
			},
		},
	}

	for _, tt := range tests {
		lexer := NewLexer(strings.NewReader(tt.input))
		for i, expected := range tt.expectedTokens {
			pos, token, literal := lexer.Lex()
			if token != expected.token {
				t.Errorf("test[%d]: expected token %s, got %s", i, expected.token, token)
			}
			if literal != expected.literal {
				t.Errorf("test[%d]: expected literal %q, got %q", i, expected.literal, literal)
			}
			if pos.Line != 1 {
				t.Errorf("test[%d]: expected line 1, got line %d", i, pos.Line)
			}
		}
	}
}
