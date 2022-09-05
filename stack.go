package evaluator

import "errors"

type stack struct {
	tokens []Token
}

func newStack(size int) *stack {
	return &stack{
		tokens: make([]Token, 0, size),
	}
}

func (s *stack) Push(t Token) error {
	if s.isFull() {
		return errors.New("stack is full")
	}
	s.tokens = append(s.tokens, t)
	return nil
}

func (s *stack) Pop() (Token, error) {
	if s.isEmpty() {
		return Token{}, errors.New("stack is empty")
	}
	token := s.tokens[len(s.tokens)-1]
	s.tokens = s.tokens[:len(s.tokens)-1]
	return token, nil
}

func (s *stack) Peek() (Token, error) {
	if s.isEmpty() {
		return Token{}, errors.New("stack is empty")
	}
	return s.tokens[len(s.tokens)-1], nil
}

func (s *stack) Size() int {
	return len(s.tokens)
}

func (s *stack) isFull() bool {
	return len(s.tokens) == cap(s.tokens)
}

func (s *stack) isEmpty() bool {
	return len(s.tokens) == 0
}
