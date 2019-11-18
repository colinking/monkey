package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	// Special chars
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	// Literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACK    = "{"
	RBRACK    = "}"

	// Keywords
	FUNC = "FUNC"
	LET  = "LET"
)
