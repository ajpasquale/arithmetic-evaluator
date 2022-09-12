package evaluator

import (
	"reflect"
	"testing"
)

func TestTokenizeInfix(t *testing.T) {
	tests := []struct {
		in   string
		want []token
	}{
		{"*", []token{{tokenOperatorMulti, "*"}}},
		{"1+1", []token{{tokenNumeric, "1"}, {tokenOperatorAdd, "+"}, {tokenNumeric, "1"}}},
		{"1+1-1", []token{{tokenNumeric, "1"}, {tokenOperatorAdd, "+"}, {tokenNumeric, "1"}, {tokenOperatorSub, "-"},
			{tokenNumeric, "1"}}},
		{"1+1-1*1", []token{{tokenNumeric, "1"}, {tokenOperatorAdd, "+"}, {tokenNumeric, "1"}, {tokenOperatorSub, "-"},
			{tokenNumeric, "1"}, {tokenOperatorMulti, "*"}, {tokenNumeric, "1"}}},
		{"1+1-1*1/1", []token{{tokenNumeric, "1"}, {tokenOperatorAdd, "+"}, {tokenNumeric, "1"}, {tokenOperatorSub, "-"},
			{tokenNumeric, "1"}, {tokenOperatorMulti, "*"}, {tokenNumeric, "1"}, {tokenOperatorDiv, "/"}, {tokenNumeric, "1"}}},
	}

	for _, tt := range tests {
		have, _ := tokenizeInfix(tt.in)

		if !reflect.DeepEqual(have, tt.want) {
			t.Errorf("tokenizer(%q)\nhave %v \nwant %v", tt.in, have, tt.want)
		}
	}
}

func TestTokenizeInfixError(t *testing.T) {
	tests := []struct {
		in   string
		want []token
	}{
		{"", []token{{tokenError, ""}}},
		{"a", []token{{tokenError, "a"}}},
		{"a+a", []token{{tokenError, "a"}}},
		{"a+1", []token{{tokenError, "a"}}},
		{"1+a", []token{{tokenError, "a"}}},
	}

	for _, tt := range tests {
		have, err := tokenizeInfix(tt.in)

		if !reflect.DeepEqual(have, tt.want) {

			t.Errorf("tokenizer(%q)\nhave %v \nwant %v", tt.in, have, tt.want)
		}
		if err == nil {
			t.Errorf("tokenizer(%q)\nhave %v \nwant tokenizer error", tt.in, have)
		}
	}
}
