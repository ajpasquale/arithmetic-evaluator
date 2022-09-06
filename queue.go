package evaluator

import (
	"errors"
)

type queue struct {
	tokens []Token
}

func newQueue() *queue {
	return &queue{
		tokens: make([]Token, 0),
	}
}

func (q *queue) Enqueue(s string) error {
	t, err := TokenizeString(s)
	if err != nil {
		return err
	}
	q.tokens = append(q.tokens, t)
	return nil
}

func (q *queue) Dequeue() (Token, error) {
	if q.isEmpty() {
		return Token{}, errors.New("the queue is empty")
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
		t, _ := tq.Dequeue()
		expr += t.val
	}
	return expr
}
