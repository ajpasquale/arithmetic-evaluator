package evaluator

import (
	"fmt"
)

func Evaluate(s string) (string, error) {
	q, _ := parse(s)
	st := newStack(10)
	for !q.isEmpty() {

		token, err := q.dequeue()
		if err != nil {
			fmt.Print(err)
		}
		switch token.typ {
		case tokenNumeric:
			st.push(token)
		case tokenOperatorAdd:
			second, _ := st.pop()
			first, _ := st.pop()
			eval, _ := computeTwoTokens(first, second, tokenOperatorAdd)
			st.push(*eval)
		case tokenOperatorSub:
			second, _ := st.pop()
			first, _ := st.pop()
			eval, _ := computeTwoTokens(first, second, tokenOperatorSub)
			st.push(*eval)
		case tokenOperatorMulti:
			second, _ := st.pop()
			first, _ := st.pop()
			eval, _ := computeTwoTokens(first, second, tokenOperatorMulti)
			st.push(*eval)
		case tokenOperatorDiv:
			second, _ := st.pop()
			first, _ := st.pop()
			eval, _ := computeTwoTokens(first, second, tokenOperatorDiv)
			st.push(*eval)
		}
	}
	sum, _ := st.pop()
	return sum.val, nil
}
