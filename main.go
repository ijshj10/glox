package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ijshj10/glox/lex"
)

func main() {
	if len(os.Args) == 1 {
		runPrompt()
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		fmt.Println("Usage: go run main.go [file]")
	}
}

func runPrompt() {
	var input string
	stdinReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		_, in_err := stdinReader.Read()
		if in_err != nil {
			fmt.Println(in_err)
			continue
		}
		err := run([]byte(input))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func runFile(filename string) {
	text, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	run(text)
}

func run(text []byte) error {
	tokens, err := lex.Lex(text)
	for _, token := range tokens {
		fmt.Println(token.Type, string(token.Lexeme))
	}
	return err
}
