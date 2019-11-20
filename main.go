package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/colinking/monkey/internal/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("👋 %s! Welcome to the Monkey REPL.\n", u.Username)
	repl.Start(os.Stdin, os.Stdout)
}
