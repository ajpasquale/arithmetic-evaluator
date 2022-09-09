package evaluator

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in   []string
		want string
	}{
		{[]string{""}, ""},
		{[]string{"1"}, "1"},
		{[]string{"+"}, "+"},
		{[]string{"1", "+", "1", "+", "1"}, "11+1+"},
		{[]string{"10", "-", "2", "*", "3", "/", "1"}, "1023*1/-"},
		{[]string{"5", "+", "3", "*", "1", "-", "9"}, "531*+9-"},
	}

	for _, tt := range tests {
		input := newQueue()
		for _, ss := range tt.in {
			input.enqueue(ss)
		}
		output, err := parse(*input)
		if err != nil {
			t.Fatal(err)
		}
		if output.String() != tt.want {
			t.Errorf("evaluator(%q)\nhave %v \nwant %v", tt.in, output.String(), tt.want)
		}
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		in []string
	}{
		{[]string{"a"}},
	}

	for _, tt := range tests {
		input := newQueue()
		for _, ss := range tt.in {
			input.enqueue(ss)
		}
		output, err := parse(*input)
		if err == nil {
			t.Errorf("evaluator(%q)\nhave %v \nwant evaluation error", tt.in, output)
		}
	}
}
