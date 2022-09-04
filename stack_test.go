package evaluator

import (
	"testing"
)

func TestEntireStack(t *testing.T) {
	s := newStack(5)

	// testing the stack creation
	if s == nil {
		t.Fatal("should have returned a new stack")
	}

	// testing the push
	for i := 1; i < 6; i++ {
		if err := s.Push(i); err != nil {
			t.Fatal("stack should not be full")
		}
	}

	// testing the stack size
	if s.Size() != 5 {
		t.Fatal("stack should be size of five")
	}

	if err := s.Push(6); err == nil {
		t.Fatal("stack should be full")
	}

	// testing the peek
	if item, _ := s.Peek(); item != 5 {
		t.Fatal("top element should be five")
	}

	// testing the pop
	for i := 5; i > 0; i-- {
		if item, _ := s.Pop(); item != i {
			t.Fatal("element should be ", i)
		}
	}

	// testing the stack size
	if s.Size() != 0 {
		t.Fatal("stack should be size of five")
	}
	if _, err := s.Pop(); err == nil {
		t.Fatal("stack should be empty")
	}
}
