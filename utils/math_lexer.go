package utils

import (
	"bufio"
	"io"
)

type Token int

const (
	EOF = iota

	MATH // $$
)

var tokens = []string{
	EOF:  "EOF",
	MATH: "$$",
}

func (t Token) String() string {
	return tokens[t]
}

type Lexer struct {
	prev   rune
	start  int
	pos    int
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (int, int, Token) {
	// keep looping until we return a token
	defer func() {
		l.start = l.pos
	}()
	out := ""
	for {
		// …
		// update the column to the position of the newly read in rune

		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.start, l.pos, EOF
			}

			// at this point there isn't much we can do, and the compiler
			// should just return the raw error to the user
			panic(err)
		}

		switch r {
		case '$':
			{
				if l.prev == 0 {
					out = out + string(r)
					l.prev = r
					l.pos++
					continue
				} else {
					if l.prev == '$' {
						l.prev = r
						l.pos++
						return l.start, l.pos - 2, MATH
					} else {
						l.prev = r
						l.pos++
					}
				}
			}
		default:
			out = out + string(r)
			l.prev = r
			l.pos++
		}
	}
}
