package parser

import (
	"fmt"
	"strings"

	"github.com/colinking/monkey/internal/lexer"
	"github.com/colinking/monkey/internal/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token

	errors []error
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  lexer,
		errors: make([]error, 0),
	}

	// Initialize curToken + peekToken.
	p.advanceToken()
	p.advanceToken()

	return p
}

func (p *Parser) Parse() *Program {
	prog := &Program{
		statements: make([]statement, 0),
	}

	for !p.curTokenIs(token.EOF) {
		var stmt statement
		switch p.curToken.Type {
		case token.Let:
			stmt = p.parseLetStatement()
		case token.Return:
			stmt = p.parseReturnStatement()
		default:
			stmt = p.parseExpressionStatement()
		}

		if stmt != nil {
			prog.statements = append(prog.statements, stmt)
		}

		if p.curTokenIs(token.Semicolon) {
			p.advanceToken()
		}
	}

	return prog
}

// advanceToken updates the parser's current and peek tokens.
func (p *Parser) advanceToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) curTokenIs(tt token.TokenType) bool {
	return p.curToken.Type == tt
}

func (p *Parser) peekTokenIs(tt token.TokenType) bool {
	return p.peekToken.Type == tt
}

func (p *Parser) expectCurTokenIs(tt token.TokenType) bool {
	if p.curToken.Type != tt {
		p.recordError(fmt.Errorf("unexpected token: '%s', expected: '%s'", p.curToken.Type, tt))
		return false
	}

	return true
}

func (p *Parser) recordError(err error) {
	p.errors = append(p.errors, err)
}

func (p *Parser) parseToken(tt token.TokenType) (token.Token, bool) {
	defer p.advanceToken()

	if !p.expectCurTokenIs(tt) {
		return token.Token{}, false
	}

	return p.curToken, true
}

// -----------
// - Program -
// -----------

// Program represents a parsed Monkey program.
type Program struct {
	statements []statement
}

func (prog *Program) String() string {
	lines := make([]string, len(prog.statements))

	for i, stmt := range prog.statements {
		lines[i] = stmt.String()
	}

	return strings.Join(lines, "\n")
}

// ---------------
// - Expressions -
// ---------------

type expression interface {
	fmt.Stringer
}

func (p *Parser) parseExpression() expression {
	for !p.curTokenIs(token.EOF) && !p.curTokenIs(token.Semicolon) {
		// TODO: pratt parser to handle:
		// integers, booleans, identifier
		// math, operator precedence, comparison
		// functions, if/else

		p.advanceToken()
	}

	return nil
}

// --------------
// - Statements -
// --------------

type statement interface {
	fmt.Stringer
}

// Let Statements

type letStatement struct {
	identifier token.Token
	value      expression
}

var _ statement = (*letStatement)(nil)

func (p *Parser) parseLetStatement() *letStatement {
	var let letStatement
	var ok bool

	if _, ok = p.parseToken(token.Let); !ok {
		return nil
	}

	if let.identifier, ok = p.parseToken(token.Identifer); !ok {
		return nil
	}

	if _, ok = p.parseToken(token.Assign); !ok {
		return nil
	}

	let.value = p.parseExpression()

	return &let
}

func (l *letStatement) String() string {
	return fmt.Sprintf("let %s = %s;", l.identifier.Literal, l.value.String())
}

// Return Statements

type returnStatement struct {
	value expression
}

var _ statement = (*returnStatement)(nil)

func (p *Parser) parseReturnStatement() *returnStatement {
	var ret returnStatement
	var ok bool

	if _, ok = p.parseToken(token.Return); !ok {
		return nil
	}

	ret.value = p.parseExpression()

	return &ret
}

func (r *returnStatement) String() string {
	return fmt.Sprintf("return %s;", r.value.String())
}

// Return Statements

type expressionStatement struct {
	value expression
}

var _ statement = (*expressionStatement)(nil)

func (e *expressionStatement) String() string {
	return fmt.Sprintf("%s;", e.value.String())
}

func (p *Parser) parseExpressionStatement() *expressionStatement {
	return &expressionStatement{
		value: p.parseExpression(),
	}
}
