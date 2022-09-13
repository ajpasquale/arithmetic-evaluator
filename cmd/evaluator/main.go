package main

import (
	"fmt"

	evaluator "github.com/ajpasquale/arithmetic-evaluator"
)

func main() {
	e, err := evaluator.Evaluate("133+233")
	if err != nil {
		fmt.Printf("error occured")
	}
	fmt.Println(e)
}
