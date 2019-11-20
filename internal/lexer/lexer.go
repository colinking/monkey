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
	ch := l.nextChar()

	// Consume any whitespace
	for isWhitespace(ch) {
		ch = l.nextChar()
	}

	if ch == 0 {
		// We've reached the end of input.
		return token.Token{Type: token.EOF}
	} else if t, ok := token.LookupChar(ch); ok {
		// Handle two-char characters.
		next := l.peekChar()
		if ch == '!' && next == '=' {
			l.position++
			return token.Token{Type: token.NotEqual, Literal: "!="}
		} else if ch == '=' && next == '=' {
			l.position++
			return token.Token{Type: token.Equal, Literal: "=="}
		} else if ch == '>' && next == '=' {
			l.position++
			return token.Token{Type: token.GreaterEqual, Literal: ">="}
		} else if ch == '<' && next == '=' {
			l.position++
			return token.Token{Type: token.LessEqual, Literal: "<="}
		}

		return token.Token{Type: t, Literal: string(ch)}
	} else if isLetter(ch) {
		// Tokenize the following characters as an identifier or keyword.
		identifier := string(ch)
		for isLetter(l.peekChar()) {
			identifier += string(l.nextChar())
		}

		return toToken(token.LookupIdentifier(identifier), identifier)
	} else if isDigit(ch) {
		// Tokenize the following characters as an integer.
		n := string(ch)
		for isDigit(l.peekChar()) {
			n += string(l.nextChar())
		}

		return toToken(token.Int, n)
	}

	// If nothing else, then this is an illegal character.
	return toToken(token.Illegal, string(ch))
}

func (l *Lexer) peekChar() byte {
	if l.position >= len(l.in) {
		return 0
	}

	return l.in[l.position]
}

func (l *Lexer) nextChar() byte {
	ch := l.peekChar()

	if ch != 0 {
		l.position++
	}

	return ch
}

func toToken(typ token.TokenType, ch string) token.Token {
	return token.Token{Type: typ, Literal: ch}
}

func isLetter(ch byte) bool {
	return ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
