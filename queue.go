package evaluator

import (
	"errors"
)

type queue struct {
	tokens []token
}

func newQueue() *queue {
	return &queue{
		tokens: make([]token, 0),
	}
}

func (q *queue) enqueue(ts ...token) {
	q.tokens = append(q.tokens, ts...)
}

func (q *queue) dequeue() (token, error) {
	if q.isEmpty() {
		return token{}, errors.New("the queue is empty")
	}
	t := q.tokens[0]
	q.tokens = q.tokens[1:]
	return t, nil
}

func (q *queue) isEmpty() bool {
	return len(q.tokens) == 0
}

func (q queue) String() string {
	tq := q
	expr := ""
	for !tq.isEmpty() {
		t, _ := tq.dequeue()
		expr += t.val
	}
	return expr
}
