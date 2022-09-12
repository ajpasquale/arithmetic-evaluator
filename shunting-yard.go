package evaluator

import (
	"errors"
)

func parse(infix string) (*queue, error) {
	ts, err := tokenizeInfix(infix)
	if err != nil {
		return nil, err
	}
	input := newQueue()
	input.enqueue(ts...)
	l := len(input.tokens)
	opStack := newStack(l)
	output := newQueue()

	for !input.isEmpty() {
		token, err := input.dequeue()
		if err != nil {
			return input, err
		}
		switch token.typ {
		case tokenNumeric:
			output.enqueue(token)
		case tokenOperatorAdd:
			fallthrough
		case tokenOperatorSub:
			fallthrough
		case tokenOperatorMulti:
			fallthrough
		case tokenOperatorDiv:
			if !opStack.isEmpty() {
				top, _ := opStack.peek()
				for precedence(top.typ) >= precedence(token.typ) {
					output.enqueue(top)
					opStack.pop()
					top, _ = opStack.peek()
				}
				opStack.push(token)
			} else {
				opStack.push(token)
			}
		case tokenError:
			return output, errors.New("encountered an error token")
		}
	}
	for !opStack.isEmpty() {
		lastToken, _ := opStack.pop()
		output.enqueue(lastToken)
	}
	return output, nil
}
