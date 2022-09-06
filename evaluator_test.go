package evaluator

import (
	"testing"
)

func TestEvaluatorEvaluate(t *testing.T) {
	q := newQueue()
	q.Enqueue("5")
	q.Enqueue("3")
	q.Enqueue("1")
	q.Enqueue("*")
	q.Enqueue("+")
	q.Enqueue("9")
	q.Enqueue("-")

	if v, _ := evaluate(q); v.val != "-1" {
		t.Fatal("evaluate should return 25")
	}
}
