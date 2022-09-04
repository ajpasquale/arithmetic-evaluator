package evaluator

import (
	"fmt"
	"strconv"
)

// evaluate "54*32*+1-"
// return "25"

// rules
/*
Numbers -> Push
Operators -> Pop (+, -, *, /)
Pop -> 2nd
Pop -> 1st
e.g. 1st operator 2nd
Push -> result
*/

func evaluate(q *pf_queue) (int, error) {
	s := newStack(10)

	for !q.isEmpty() {
		token, err := q.Dequeue()
		if err != nil {
			fmt.Print(err)
		}
		switch token.typ {
		case tokenNumeric:
			v, _ := strconv.Atoi(token.val)
			s.Push(v)
		case tokenOperatorAdd:
			second, _ := s.Pop()
			first, _ := s.Pop()

			eval := first + second

			s.Push(eval)

		case tokenOperatorSub:
			second, _ := s.Pop()
			first, _ := s.Pop()

			eval := first - second

			s.Push(eval)
		case tokenOperatorMulti:
			second, _ := s.Pop()
			first, _ := s.Pop()

			eval := first * second

			s.Push(eval)
		case tokenOperatorDiv:
			second, _ := s.Pop()
			first, _ := s.Pop()

			eval := first / second

			s.Push(eval)
		}
	}
	return s.Pop()
}
