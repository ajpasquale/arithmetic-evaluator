package evaluator

import (
	"fmt"
)

func Evaluate(s string) (string, error) {
	q, _ := tokenizeInfix(s)
	q, _ = parse(q)
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
			eval, _ := addTwoTokens(first, second)
			st.push(eval)
		case tokenOperatorSub:
			second, _ := st.pop()
			first, _ := st.pop()
			eval, _ := subTwoTokens(first, second)
			st.push(eval)
		case tokenOperatorMulti:
			second, _ := st.pop()
			first, _ := st.pop()
			eval, _ := multiplyTwoTokens(first, second)
			st.push(eval)
		case tokenOperatorDiv:
			second, _ := st.pop()
			first, _ := st.pop()
			eval, _ := dividTwoTokens(first, second)
			st.push(eval)
		}
	}
	sum, _ := st.pop()
	return sum.val, nil
}
