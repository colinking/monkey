package lexer

import (
	"github.com/colinking/monkey/internal/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextToken(t *testing.T) {
	in := "=+,;(){}"

	tokens := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.PLUS, expectedLiteral: "+"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.LBRACK, expectedLiteral: "{"},
		{expectedType: token.RBRACK, expectedLiteral: "}"},
	}

	l := New(in)

	for _, expected := range tokens {
		token := l.NextToken()
		require.Equal(t, expected.expectedType, token.Type)
		require.Equal(t, expected.expectedLiteral, token.Literal)
	}
}
