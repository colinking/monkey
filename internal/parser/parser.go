package parser

import (
	"github.com/colinking/monkey/internal/lexer"
	"github.com/colinking/monkey/internal/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: lexer,
	}

	// Initialize curToken + peekToken.
	p.advanceToken()
	p.advanceToken()

	return p
}

func (p *Parser) Parse() *Program {
	prog := &Program{}

	for p.curToken.Type != token.EOF {
		switch p.curToken.Type {
		case token.Let:
			// TODO
		case token.Return:
			// TODO
		default:
			// TODO
		}
	}

	return prog
}

// advanceToken updates the parser's current and peek tokens.
func (p *Parser) advanceToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// -----------
// - Program -
// -----------

// Program represents a parsed Monkey program.
type Program struct {
	statements *[]statement
}

// --------------
// - Statements -
// --------------

type statement interface {
	// TODO
}

// Let

type letStatement struct {
	// TODO
}

// Return

type returnStatement struct {
	// TODO
}

// ---------------
// - Expressions -
// ---------------

type expression interface {
	// TODO
}
