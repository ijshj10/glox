package lex

import (
	"fmt"
)

type Type string

type Token struct {
	Type   Type
	Lexeme []byte
	Line   int
}

func Lex(text []byte) ([]Token, error) {
	line := 1
	offset := 0
	tokens := make([]Token, 0)
	for offset < len(text) {
		switch text[offset] {
		case '\n':
			line++
			offset++
		case '\t', ' ', '\r':
			offset++
		case '=':
			if offset+1 < len(text) && text[offset+1] == '=' {
				tokens = append(tokens, Token{Type: "==", Lexeme: text[offset : offset+2], Line: line})
				offset += 2
			} else {
				tokens = append(tokens, Token{Type: "=", Lexeme: text[offset : offset+1], Line: line})
				offset++
			}
		case '!':
			if offset+1 < len(text) && text[offset+1] == '=' {
				tokens = append(tokens, Token{Type: "!=", Lexeme: text[offset : offset+2], Line: line})
				offset += 2
			} else {
				tokens = append(tokens, Token{Type: "!", Lexeme: text[offset : offset+1], Line: line})
				offset++
			}
		case '+':
			tokens = append(tokens, Token{Type: "+", Lexeme: text[offset : offset+1], Line: line})
			offset++
		case '-':
			tokens = append(tokens, Token{Type: "-", Lexeme: text[offset : offset+1], Line: line})
			offset++
		case '*':
			tokens = append(tokens, Token{Type: "*", Lexeme: text[offset : offset+1], Line: line})
			offset++
		case '/':
			if offset+1 < len(text) && text[offset+1] == '/' {
				for offset < len(text) && text[offset] != '\n' {
					offset++
				}
			} else {
				tokens = append(tokens, Token{Type: "/", Lexeme: text[offset : offset+1], Line: line})
				offset++
			}
		case '<':
			if offset+1 < len(text) && text[offset+1] == '=' {
				tokens = append(tokens, Token{Type: "<=", Lexeme: text[offset : offset+2], Line: line})
				offset += 2
			} else {
				tokens = append(tokens, Token{Type: "<", Lexeme: text[offset : offset+1], Line: line})
				offset++
			}
		case '>':
			if offset+1 < len(text) && text[offset+1] == '=' {
				tokens = append(tokens, Token{Type: ">=", Lexeme: text[offset : offset+2], Line: line})
				offset += 2
			} else {
				tokens = append(tokens, Token{Type: ">", Lexeme: text[offset : offset+1], Line: line})
				offset++
			}
		case ';':
			tokens = append(tokens, Token{Type: ";", Lexeme: text[offset : offset+1], Line: line})
			offset++
		case ',':
			tokens = append(tokens, Token{Type: ",", Lexeme: text[offset : offset+1], Line: line})
			offset++
		case '(':
			tokens = append(tokens, Token{Type: "(", Lexeme: text[offset : offset+1], Line: line})
			offset++
		case ')':
			tokens = append(tokens, Token{Type: ")", Lexeme: text[offset : offset+1], Line: line})
			offset++
		case '{':
			tokens = append(tokens, Token{Type: "{", Lexeme: text[offset : offset+1], Line: line})
			offset++
		case '}':
			tokens = append(tokens, Token{Type: "}", Lexeme: text[offset : offset+1], Line: line})
			offset++
		case '"':
			lexeme := make([]byte, 0)
			start_line := line
			offset++
			for offset < len(text) && text[offset] != '"' {
				if text[offset] == '\n' {
					line++
				}
				lexeme = append(lexeme, text[offset])
				offset++
			}
			if offset >= len(text) {
				return nil, fmt.Errorf("Unterminated \" at %d", start_line)
			}
			tokens = append(tokens, Token{Type: "string", Lexeme: lexeme, Line: start_line})
			offset++
		default:
			lexeme := make([]byte, 0)
			if isNumber(text[offset]) {
				for offset < len(text) && isNumber(text[offset]) {
					lexeme = append(lexeme, text[offset])
					offset++
				}
				if offset < len(text) && text[offset] == '.' {
					lexeme = append(lexeme, text[offset])
					offset++
					for offset < len(text) && isNumber(text[offset]) {
						lexeme = append(lexeme, text[offset])
						offset++
					}
				}
				tokens = append(tokens, Token{Type: "number", Lexeme: lexeme, Line: line})
			} else if isAlpha(text[offset]) {
				for offset < len(text) && isAlpha(text[offset]) {
					lexeme = append(lexeme, text[offset])
					offset++
				}
				tokens = append(tokens, Token{Type: "identifier", Lexeme: lexeme, Line: line})
			} else {
				return nil, fmt.Errorf("Unexpected character %c at %d", text[offset], line)
			}
		}
	}
	return tokens, nil
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isNumber(c)
}

const (
	// Single-character tokens.
	LEFT_PAREN  Type = "("
	RIGHT_PAREN Type = ")"
	LEFT_BRACE  Type = "{"
	RIGHT_BRACE Type = "}"
	COMMA       Type = ","
	DOT         Type = "."
	MINUS       Type = "-"
	PLUS        Type = "+"
	SEMICOLON   Type = ";"
	SLASH       Type = "/"
	STAR        Type = "*"

	// One or two character tokens.
	BANG          Type = "!"
	BANG_EQUAL    Type = "!="
	EQUAL         Type = "="
	EQUAL_EQUAL   Type = "=="
	GREATER       Type = ">"
	GREATER_EQUAL Type = ">="
	LESS          Type = "<"

	// Literals.
	IDENTIFIER Type = "IDENTIFIER"
	STRING     Type = "STRING"
	NUMBER     Type = "NUMBER"

	// Keywords.
	AND    Type = "AND"
	CLASS  Type = "CLASS"
	ELSE   Type = "ELSE"
	FALSE  Type = "FALSE"
	FUN    Type = "FUN"
	FOR    Type = "FOR"
	IF     Type = "IF"
	NIL    Type = "NIL"
	OR     Type = "OR"
	PRINT  Type = "PRINT"
	RETURN Type = "RETURN"
	SUPER  Type = "SUPER"
	THIS   Type = "THIS"
	TRUE   Type = "TRUE"
	VAR    Type = "VAR"
	WHILE  Type = "WHILE"

	EOF Type = "EOF"
)
