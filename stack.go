package evaluator

import "errors"

type stack struct {
	items []int
}

func newStack(size int) *stack {
	return &stack{
		items: make([]int, 0, size),
	}
}

func (s *stack) Push(item int) error {
	if !s.isFull() {
		s.items = append(s.items, item)
		return nil
	}
	return errors.New("stack is full")
}

func (s *stack) Pop() (int, error) {
	if !s.isEmpty() {
		item := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return item, nil
	}
	return 0, errors.New("stack is empty")
}

func (s *stack) Peek() (int, error) {
	if !s.isEmpty() {
		return s.items[len(s.items)-1], nil
	}
	return 0, errors.New("stack is empty")
}

func (s *stack) Size() int {
	return len(s.items)
}

func (s *stack) isFull() bool {
	return len(s.items) == cap(s.items)
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}
