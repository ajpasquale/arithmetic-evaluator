package evaluator

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"1", "1"},
		{"+", "+"},
		{"1+1+1", "11+1+"},
		{"10-2*3/1", "1023*1/-"},
		{"5+3*1-9", "531*+9-"},
	}

	for _, tt := range tests {
		output, _ := parse(tt.in)
		if output.String() != tt.want {
			t.Errorf("evaluator(%q)\nhave %v \nwant %v", tt.in, output.String(), tt.want)
		}
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		in string
	}{
		{""},
		{"a"},
		{"\\"},
		{"="},
		{"a+b*c-1"},
		{"hello world"},
		{"22-a-32+1"},
	}

	for _, tt := range tests {
		output, err := parse(tt.in)
		if err == nil {
			t.Errorf("evaluator(%q)\nhave %v \nwant evaluation error", tt.in, output)
		}
	}
}
