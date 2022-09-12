package evaluator

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		in   string
		want []token
	}{
		{"1*1/1", []token{{tokenNumeric, "1"}, {tokenOperatorMulti, "*"}, {tokenNumeric, "1"}, {tokenOperatorDiv, "/"}, {tokenNumeric, "1"}}},
	}

	for _, tt := range tests {
		q := newQueue()
		if !q.isEmpty() {
			t.Fatal("queue should be empty")
		}

		s, _ := tokenizeInfix(tt.in)

		q.enqueue(s...)
		if q.isEmpty() {
			t.Fatal("queue should not be empty")
		}

		for i := 0; !q.isEmpty(); i++ {
			have, err := q.dequeue()
			if err != nil {
				t.Fatal("dequeue error")
			}

			if !reflect.DeepEqual(have, tt.want[i]) {
				t.Errorf("queue(%q)\nhave %v \nwant %v", tt.in, have, tt.want)
			}
		}

		if !q.isEmpty() {
			t.Fatal("queue should be empty")
		}
	}
}
