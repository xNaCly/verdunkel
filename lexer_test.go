package verdunkel

import (
	"strings"
	"testing"
)

func Compare(a, b []Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {

		if a[i].Type != b[i].Type {
			return false
		} else if a[i].Content != b[i].Content {
			return false
		}
	}

	return true
}

func TestLexer(t *testing.T) {
	tests := []struct {
		in  string
		out []Token
	}{
		{"package verdunkel", []Token{}},
	}
	for _, test := range tests {
		l := Lexer{}
		out, err := l.Lex(strings.NewReader(test.in))
		if err != nil {
			t.Error("failed to lex")
		}
		if !Compare(test.out, out) {
			t.Error("not the expected output")
		}
	}
}
