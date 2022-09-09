package evaluator

import (
	"testing"
)

func TestTokenizeInfix(t *testing.T) {
	q, err := tokenizeInfix("1+1+1+1")
	if err != nil {
		t.Fatal(err)
	}

	have := q.String()
	if have != "1+1+1+1" {
		t.Fatal("")
	}
}
