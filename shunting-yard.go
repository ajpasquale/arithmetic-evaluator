package evaluator

import (
	"errors"
)

func Parse(input queue) (queue, error) {
	l := len(input.tokens)
	opStack := newStack(l)
	output := newQueue()

	for !input.isEmpty() {
		token, err := input.Dequeue()
		if err != nil {
			return input, err
		}
		switch token.typ {
		case tokenNumeric:
			output.Enqueue(token.val)
		case tokenOperatorAdd:
			fallthrough
		case tokenOperatorSub:
			if !opStack.isEmpty() {
				if t, _ := opStack.Peek(); t.typ < tokenOperatorAdd {
					for !opStack.isEmpty() {
						tk, _ := opStack.Pop()
						output.Enqueue(tk.val)
					}
					opStack.Push(token)
				}
			} else {
				opStack.Push(token)
			}
		case tokenOperatorMulti:
			fallthrough
		case tokenOperatorDiv:
			opStack.Push(token)
		case tokenError:
			return *output, errors.New("encountered an error token")
		}
	}

	lastToken, _ := opStack.Pop()
	output.Enqueue(lastToken.val)

	return *output, nil
}
