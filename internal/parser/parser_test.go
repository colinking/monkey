package parser

import (
	"testing"

	"github.com/colinking/monkey/internal/lexer"
	"github.com/colinking/monkey/internal/token"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	input := `
		let foo = 123;
		let hello_world = 456;
		return 100;
		func test() {};
	`
	output := &Program{
		statements: []statement{
			&letStatement{
				identifier: token.Token{
					Type:    token.Identifer,
					Literal: "foo",
				},
				value: nil,
			},
			&letStatement{
				identifier: token.Token{
					Type:    token.Identifer,
					Literal: "hello_world",
				},
				value: nil,
			},
			&returnStatement{
				value: nil,
			},
			&expressionStatement{
				value: nil,
			},
		},
	}

	l := lexer.New(input)
	p := New(l)
	prog := p.Parse()

	require.NotNil(t, prog)
	require.ElementsMatch(t, output.statements, prog.statements)
}
