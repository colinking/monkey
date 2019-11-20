package lexer

import (
	"testing"

	"github.com/colinking/monkey/internal/token"

	"github.com/stretchr/testify/require"
)

func TestNextToken(s *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedTokens []token.Token
	}{
		{
			name: "Lexer skips whitespace",
			input: `	 
`,
			expectedTokens: []token.Token{
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Lexer tokenizes one-char tokens",
			input: "=+-!/* <> ,;(){} ~",
			expectedTokens: []token.Token{
				{Type: token.Assign, Literal: "="},
				{Type: token.Plus, Literal: "+"},
				{Type: token.Minus, Literal: "-"},
				{Type: token.Not, Literal: "!"},
				{Type: token.Slash, Literal: "/"},
				{Type: token.Mult, Literal: "*"},

				{Type: token.Less, Literal: "<"},
				{Type: token.Greater, Literal: ">"},

				{Type: token.Comma, Literal: ","},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.LParen, Literal: "("},
				{Type: token.RParen, Literal: ")"},
				{Type: token.LBrack, Literal: "{"},
				{Type: token.RBrack, Literal: "}"},

				{Type: token.Illegal, Literal: "~"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Lexer tokenizes identifiers",
			input: "foo + bar",
			expectedTokens: []token.Token{
				{Type: token.Identifer, Literal: "foo"},
				{Type: token.Plus, Literal: "+"},
				{Type: token.Identifer, Literal: "bar"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Lexer tokenizes numbers",
			input: "123 + 456",
			expectedTokens: []token.Token{
				{Type: token.Int, Literal: "123"},
				{Type: token.Plus, Literal: "+"},
				{Type: token.Int, Literal: "456"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Lexer tokenizes keywords",
			input: "func let return",
			expectedTokens: []token.Token{
				{Type: token.Func, Literal: "func"},
				{Type: token.Let, Literal: "let"},
				{Type: token.Return, Literal: "return"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "Lexer tokenizes two-char comparison operators",
			input: "== != >= <=",
			expectedTokens: []token.Token{
				{Type: token.Equal, Literal: "=="},
				{Type: token.NotEqual, Literal: "!="},
				{Type: token.GreaterEqual, Literal: ">="},
				{Type: token.LessEqual, Literal: "<="},
			},
		},
		{
			name:  "Lexer tokenizes keywords",
			input: `func let true false if else return`,
			expectedTokens: []token.Token{
				{Type: token.Func, Literal: "func"},
				{Type: token.Let, Literal: "let"},
				{Type: token.True, Literal: "true"},
				{Type: token.False, Literal: "false"},
				{Type: token.If, Literal: "if"},
				{Type: token.Else, Literal: "else"},
				{Type: token.Return, Literal: "return"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func(t *testing.T) {
			l := New(test.input)
			tokens := make([]token.Token, 0)
			for _, expected := range test.expectedTokens {
				token := l.NextToken()
				tokens = append(tokens, token)
				require.Equal(t, expected.Type, token.Type, "Unexpected token type. Tokens: %v", tokens)
				require.Equal(t, expected.Literal, token.Literal, "Unexpected token literal. Tokens: %v", tokens)
			}
		})
	}
}
