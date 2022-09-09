package evaluator

import (
	"reflect"
	"testing"
)

func TestEvaluator(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"1", "1"},
		{"1+1", "2"},
		{"1+1+1", "3"},
		{"1+1+1+1", "4"},
		{"1-1", "0"},
		{"2-1", "1"},
		{"3-1", "2"},
		{"4-2", "2"},
		{"4-3", "1"},
		{"4-4", "0"},
		{"5-1-1", "3"},
		{"5-1-1-1", "2"},
		{"1*1", "1"},
		{"2*2", "4"},
		{"2*2*2", "8"},
		{"2*2*2*2", "16"},
		{"0/1", "0"},
		{"1/1", "1"},
		{"8/2", "4"},
		{"8/2/2", "2"},
		{"5+2*3-9", "2"},
		{"10-2*3/1", "4"},
		{"2*8-2+3", "17"},
		{"82-7+10*2", "95"},
		{"2*2*4+2+1/2", "18"},
	}

	for _, tt := range tests {
		have, err := Evaluate(tt.in)
		if err != nil {
			t.Fatalf("evaluator(%q) %v", tt.in, err)
		}

		if !reflect.DeepEqual(have, tt.want) {
			t.Errorf("evaluator(%q)\nhave %v \nwant %v", tt.in, have, tt.want)
		}
	}
}
