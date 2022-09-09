package evaluator

import (
	"testing"
)

// turns an infix expression like "5+3*1-9" into a queue of tokens
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
