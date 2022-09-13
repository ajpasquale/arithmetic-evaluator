# evaluator

Package evaluator implements a shunting-yard algorithm to parse an infix expression and evaluate the postfix result

parse an expression and return the result

```go
e, err := evaluator.Evaluate("1+2*3-4/1")
```