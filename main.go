package main

import (
	"fmt"
	"os"
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
	for {
		fmt.Print("> ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}
		run([]byte(input))
	}
}

func runFile(filename string) {
	text, err := os.ReadFile(filename)
	if err != nil {
		println(err.Error())
		return
	}
	run(text)
}

func run(text []byte) error {
	tokens, err := Lex(text)
	for _, token := range tokens {
		fmt.Println(token)
	}
	return err
}

type Type string

type Token struct {
	Type   Type
	lexeme []byte
}

func Lex(text []byte) ([]Token, error) {
	return nil, nil
}
