package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/colinking/monkey/internal/lexer"
	"github.com/colinking/monkey/internal/token"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, prompt)

		if ok := scanner.Scan(); !ok {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
