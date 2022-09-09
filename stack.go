package evaluator

import "errors"

type stack struct {
	tokens []token
}

func newStack(size int) *stack {
	return &stack{
		tokens: make([]token, 0, size),
	}
}

func (s *stack) push(t token) error {
	if s.isFull() {
		return errors.New("stack is full")
	}
	s.tokens = append(s.tokens, t)
	return nil
}

func (s *stack) pop() (token, error) {
	if s.isEmpty() {
		return token{}, errors.New("stack is empty")
	}
	token := s.tokens[len(s.tokens)-1]
	s.tokens = s.tokens[:len(s.tokens)-1]
	return token, nil
}

func (s *stack) peek() (token, error) {
	if s.isEmpty() {
		return token{}, errors.New("stack is empty")
	}
	return s.tokens[len(s.tokens)-1], nil
}

func (s *stack) isFull() bool {
	return len(s.tokens) == cap(s.tokens)
}

func (s *stack) isEmpty() bool {
	return len(s.tokens) == 0
}
