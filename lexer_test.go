package graphite_test

import (
	"testing"

	"github.com/jsternberg/graphite"
)

func TestLexer(t *testing.T) {
	// This query makes no sense, but it tests nested functions and parameters
	// and we're not actually attempting to use the query for anything.
	inp := []byte("aliasByNode(sumSeries(hello.*.world), 1)")
	lexer := graphite.NewLexer(inp)

	tokens := []graphite.Token{
		{"aliasByNode", graphite.Ident},
		{"(", graphite.LeftParen},
		{"sumSeries", graphite.Ident},
		{"(", graphite.LeftParen},
		{"hello.*.world", graphite.Ident},
		{")", graphite.RightParen},
		{",", graphite.Separator},
		{"1", graphite.Number},
		{")", graphite.RightParen},
	}

	for i, tok := range tokens {
		if !lexer.Scan() {
			t.Errorf("expected %v, got nil", tok)
			break
		} else if tokens[i] != tok {
			t.Errorf("expected %v, got %v", tokens[i], tok)
			break
		}
	}

	if err := lexer.Error(); err != nil {
		t.Fatal(err)
	}

	if lexer.Scan() {
		t.Fatalf("expected eof, got %v", lexer.Token())
	}
}
