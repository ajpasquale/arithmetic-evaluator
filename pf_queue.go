package evaluator

import (
	"errors"
	"strconv"
)

type tokenType int

const (
	tokenNumeric tokenType = iota
	tokenOperatorAdd
	tokenOperatorSub
	tokenOperatorMulti
	tokenOperatorDiv
)

type Token struct {
	typ tokenType
	val string
}

type pf_queue struct {
	tokens []Token
}

func newQueue() *pf_queue {
	return &pf_queue{
		tokens: make([]Token, 0),
	}
}

func (q *pf_queue) Enqueue(s string) error {
	switch {
	case isNumeric(s):
		q.tokens = append(q.tokens, Token{tokenNumeric, s})
		return nil
	case s == "+":
		q.tokens = append(q.tokens, Token{tokenOperatorAdd, s})
		return nil
	case s == "-":
		q.tokens = append(q.tokens, Token{tokenOperatorSub, s})
		return nil
	case s == "*":
		q.tokens = append(q.tokens, Token{tokenOperatorMulti, s})
		return nil
	case s == "/":
		q.tokens = append(q.tokens, Token{tokenOperatorDiv, s})
		return nil
	}
	return errors.New("invalid token")
}

func (q *pf_queue) Dequeue() (Token, error) {
	if q.isEmpty() {
		return Token{}, errors.New("the queue is empty")
	}
	t := q.tokens[0]
	q.tokens = q.tokens[1:]
	return t, nil
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func (q *pf_queue) isEmpty() bool {
	return len(q.tokens) == 0
}
