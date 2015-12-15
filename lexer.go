package graphite

import "fmt"

type TokenType int

const (
	None       TokenType = 0
	Ident                = iota
	LeftParen            = iota
	RightParen           = iota
	Separator            = iota
	Number               = iota
)

type Token struct {
	Value string
	Type  TokenType
}

type Lexer struct {
	p   int
	pe  int
	eof int

	data []byte
	cs   int
	ts   int
	te   int
	act  int

	lineno  int
	linebeg int

	token Token
	err   error
}

func NewLexer(data []byte) *Lexer {
	l := &Lexer{}
	l.lineno = 1
	l.init(data)
	return l
}

func (l *Lexer) Scan() bool {
	if l.p == l.eof {
		return false
	}
	l.clear()
	l.scan()

	if l.cs == graphite_error {
		l.err = fmt.Errorf("unexpected token '%c' at line %d, col %d", l.data[l.p], l.lineno, l.p-l.linebeg)
		return false
	} else if l.token.Type == None {
		return false
	}
	return true
}

func (l *Lexer) Token() Token {
	return l.token
}

func (l *Lexer) Error() error {
	return l.err
}

func (l *Lexer) clear() {
	l.token = Token{}
}

func (l *Lexer) emit(typ TokenType) {
	value := string(l.data[l.ts:l.te])
	l.token = Token{
		Value: value,
		Type:  typ,
	}
}
