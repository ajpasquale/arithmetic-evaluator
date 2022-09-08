package evaluator

import (
	"fmt"
)

func evaluate(q *queue) (Token, error) {
	s := newStack(10)

	for !q.isEmpty() {
		token, err := q.Dequeue()
		if err != nil {
			fmt.Print(err)
		}
		switch token.typ {
		case tokenNumeric:
			s.Push(token)
		case tokenOperatorAdd:
			second, _ := s.Pop()
			first, _ := s.Pop()
			eval, _ := AddTwoTokens(first, second)
			s.Push(eval)
		case tokenOperatorSub:
			second, _ := s.Pop()
			first, _ := s.Pop()
			eval, _ := SubTwoTokens(first, second)
			s.Push(eval)
		case tokenOperatorMulti:
			second, _ := s.Pop()
			first, _ := s.Pop()
			eval, _ := MultiplyTwoTokens(first, second)
			s.Push(eval)
		case tokenOperatorDiv:
			second, _ := s.Pop()
			first, _ := s.Pop()
			eval, _ := DividTwoTokens(first, second)
			s.Push(eval)
		}
	}
	return s.Pop()
}
