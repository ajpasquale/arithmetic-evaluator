package evaluator

import (
	"fmt"
	"testing"
)

func TestEvaluatorEvaluate(t *testing.T) {
	q := newQueue()
	q.Enqueue("5")
	q.Enqueue("4")
	q.Enqueue("*")
	q.Enqueue("3")
	q.Enqueue("2")
	q.Enqueue("*")
	q.Enqueue("+")
	q.Enqueue("1")
	q.Enqueue("-")

	v, _ := evaluate(q)
	fmt.Print(v)

}
