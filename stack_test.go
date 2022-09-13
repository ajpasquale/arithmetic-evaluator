package evaluator

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		in   string
		want []token
	}{
		{"1*1/1", []token{{tokenNumeric, "1"}, {tokenOperatorMulti, "*"}, {tokenNumeric, "1"}, {tokenOperatorDiv, "/"}, {tokenNumeric, "1"}}},
	}

	for _, tt := range tests {
		stk := newStack(len(tt.in))

		if !stk.isEmpty() {
			t.Fatal("stack should be empty")
		}

		tks, _ := tokenizeInfix(tt.in)

		stk.push(tks...)
		if stk.isEmpty() {
			t.Fatal("stack should not be empty")
		}

		for i := len(tt.want) - 1; !stk.isEmpty(); i-- {
			have, err := stk.pop()
			if err != nil {
				t.Fatal("pop error")
			}

			if !reflect.DeepEqual(have, tt.want[i]) {
				t.Errorf("stack(%q)\nhave %v \nwant %v", tt.in, have, tt.want)
			}
		}

		if !stk.isEmpty() {
			t.Fatal("stack should be empty")
		}
	}
}
