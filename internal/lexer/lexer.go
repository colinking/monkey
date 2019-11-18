package lexer

import (
	"github.com/colinking/monkey/internal/token"
)

type Lexer struct {
	in       string
	position int
}

// New constructs a new lexer instance that will tokenize the provided string.
func New(in string) *Lexer {
	return &Lexer{
		in: in,
	}
}

// NextToken returns the next token using this lexer.
func (l *Lexer) NextToken() token.Token {
	// TODO

	return token.Token{}
}
