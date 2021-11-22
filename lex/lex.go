package lex

type Type string

type Token struct {
	Type   Type
	lexeme []byte
}

func Lex(text []byte) ([]Token, error) {
	return nil, nil
}
