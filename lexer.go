package verdunkel

import (
	"bufio"
	"io"
)

type TokenType uint8

const (
	EOF TokenType = iota
	CONSTANT
	IDENTIFIER
	TYPE
)

type Token struct {
	Type    TokenType
	Content string
}

type Lexer struct {
	input []rune
	i     int
}

func (l *Lexer) Lex(r io.Reader) ([]Token, error) {
	b := bufio.NewReader(r)
	l.input = make([]rune, 0, b.Size())
	// consume whole input, make indexable as array
	for {
		r, _, err := b.ReadRune()
		if err != nil {
			break
		}
		l.input = append(l.input, r)
	}
	l.i = 0
	return nil, nil
}

func (l *Lexer) makeToken(t TokenType, length int) Token {
	return Token{Type: t, Content: string(l.input[l.i-length : l.i])}
}
