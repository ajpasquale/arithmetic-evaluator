package evaluator

import (
	"errors"
	"strconv"
	"strings"
)

type tokenType int

const (
	tokenError tokenType = iota
	tokenOperatorParens
	tokenOperatorExpo
	tokenOperatorMulti
	tokenOperatorDiv
	tokenOperatorAdd
	tokenOperatorSub
	tokenNumeric
)

type token struct {
	typ tokenType
	val string
}

func tokenizeInt(i int) *token {
	s := strconv.Itoa(i)
	return &token{
		tokenNumeric,
		s,
	}
}

func tokenizeInfix(s string) ([]token, error) {
	if len(s) == 0 {
		t, err := tokenizeString(s)
		return []token{t}, err
	}
	var substrs []string
	for {
		i := strings.IndexFunc(s, func(r rune) bool {
			return r == '+' || r == '-' || r == '*' || r == '/'
		})
		if i == -1 || len(s) == 1 {
			if len(s) > 0 {
				substrs = append(substrs, s)
			}
			break
		}
		substrs = append(substrs, strings.TrimSpace(s[:i]), s[i:i+1])
		s = strings.TrimSpace(s[i+1:])
	}

	tokens := make([]token, 0)
	for _, v := range substrs {
		t, err := tokenizeString(v)
		if err != nil {
			return []token{t}, err
		}
		tokens = append(tokens, t)
	}
	return tokens, nil
}

func tokenizeString(s string) (token, error) {
	switch {
	case isNumeric(s):
		return token{tokenNumeric, s}, nil
	case s == "+":
		return token{tokenOperatorAdd, s}, nil
	case s == "-":
		return token{tokenOperatorSub, s}, nil
	case s == "*":
		return token{tokenOperatorMulti, s}, nil
	case s == "/":
		return token{tokenOperatorDiv, s}, nil
	}
	return token{tokenError, s}, errors.New("tokenizeString: unknown token")
}

func computeTwoTokens(f token, s token, op tokenType) (*token, error) {
	if f.typ == tokenNumeric && s.typ == tokenNumeric {
		fi, _ := strconv.Atoi(f.val)
		si, _ := strconv.Atoi(s.val)

		ni := sumTwoValue(fi, si, op)

		return tokenizeInt(ni), nil
	}
	return nil, errors.New("computeTwoTokens: cannot compute non-numeric tokens")
}
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func Split(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}

func sumTwoValue(f int, s int, op tokenType) int {
	switch op {
	case tokenOperatorAdd:
		return f + s
	case tokenOperatorSub:
		return f - s
	case tokenOperatorMulti:
		return f * s
	case tokenOperatorDiv:
		return f / s
	}
	return 0
}

func precedence(typ tokenType) int {
	switch typ {
	case tokenOperatorAdd:
		fallthrough
	case tokenOperatorSub:
		return 1
	case tokenOperatorMulti:
		fallthrough
	case tokenOperatorDiv:
		return 2
	}
	return 0
}
