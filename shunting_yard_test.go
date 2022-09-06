package evaluator

import (
	"testing"
)

func TestParse(t *testing.T) {
	q := newQueue()
	q.Enqueue("5")
	q.Enqueue("+")
	q.Enqueue("3")
	q.Enqueue("*")

	q.Enqueue("1")
	q.Enqueue("-")
	q.Enqueue("9")
	expr := q.String()
	if expr != "5+3*1-9" {
		t.Fatal("String method did not return the correct value")
	}
	o, _ := Parse(*q)
	if o.String() != "531*+9-" {
		t.Fatal("String method did not return the correct value")
	}
}
