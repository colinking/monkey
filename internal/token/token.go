package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	// Special chars
	Illegal TokenType = "illegal"
	EOF               = "eof"

	// Literals
	Identifer = "identifer"
	Int       = "int"

	// Operators
	Assign = "="
	Plus   = "+"
	Minus  = "-"
	Not    = "!"
	Slash  = "/"
	Mult   = "*"

	// Comparison
	Equal        = "=="
	NotEqual     = "!="
	Less         = "<"
	LessEqual    = "<="
	Greater      = ">"
	GreaterEqual = ">="

	// Delimiters
	Comma     = ","
	Semicolon = ";"
	LParen    = "("
	RParen    = ")"
	LBrack    = "{"
	RBrack    = "}"

	// Keywords
	Func   = "func"
	Let    = "let"
	True   = "true"
	False  = "false"
	If     = "if"
	Else   = "else"
	Return = "return"
)

var keywords = map[string]TokenType{
	"func":   Func,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookupIdentifier(identifier string) TokenType {
	if t, ok := keywords[identifier]; ok {
		return t
	}

	return Identifer
}

var specialChars = map[byte]TokenType{
	'=': Assign,
	'+': Plus,
	'-': Minus,
	'!': Not,
	'/': Slash,
	'*': Mult,
	'<': Less,
	'>': Greater,
	',': Comma,
	';': Semicolon,
	'(': LParen,
	')': RParen,
	'{': LBrack,
	'}': RBrack,
}

func LookupChar(ch byte) (TokenType, bool) {
	t, ok := specialChars[ch]
	return t, ok
}
