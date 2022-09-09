package evaluator

import (
	"errors"
)

func parse(input queue) (queue, error) {
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
			output.enqueue(token.val)
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
					output.enqueue(top.val)
					opStack.pop()
					top, _ = opStack.peek()
				}
				opStack.push(token)
			} else {
				opStack.push(token)
			}
		case tokenError:
			return *output, errors.New("encountered an error token")
		}
	}
	for !opStack.isEmpty() {
		lastToken, _ := opStack.pop()
		output.enqueue(lastToken.val)
	}
	return *output, nil
}
