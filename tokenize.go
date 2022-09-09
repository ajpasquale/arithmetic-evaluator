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

// turns an infix expression like "5+3*1-9" into a queue of tokens
func tokenizeInfix(s string) (queue, error) {
	q := newQueue()

	var substrs []string
	for {
		i := strings.IndexFunc(s, func(r rune) bool {
			return r == '+' || r == '-' || r == '*' || r == '/'
		})
		if i == -1 {
			if len(s) > 0 {
				substrs = append(substrs, s)
			}
			break
		}
		substrs = append(substrs, strings.TrimSpace(s[:i]), s[i:i+1])
		s = strings.TrimSpace(s[i+1:])
	}

	for _, v := range substrs {
		q.enqueue(v)
	}
	return *q, nil
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
	return token{tokenError, ""}, errors.New("unknown token")
}

func addTwoTokens(f token, s token) (token, error) {
	if f.typ == tokenNumeric && s.typ == tokenNumeric {
		fi, _ := strconv.Atoi(f.val)
		si, _ := strconv.Atoi(s.val)

		ni := fi + si

		return *tokenizeInt(ni), nil
	}
	return token{}, errors.New("cannot add non-numeric tokens")
}

func subTwoTokens(f token, s token) (token, error) {
	if f.typ == tokenNumeric && s.typ == tokenNumeric {
		fi, _ := strconv.Atoi(f.val)
		si, _ := strconv.Atoi(s.val)

		ni := fi - si

		return *tokenizeInt(ni), nil
	}
	return token{}, errors.New("cannot subtract non-numeric tokens")
}

func multiplyTwoTokens(f token, s token) (token, error) {
	if f.typ == tokenNumeric && s.typ == tokenNumeric {
		fi, _ := strconv.Atoi(f.val)
		si, _ := strconv.Atoi(s.val)

		ni := fi * si

		return *tokenizeInt(ni), nil
	}
	return token{}, errors.New("cannot multiple non-numeric tokens")
}

func dividTwoTokens(f token, s token) (token, error) {
	if f.typ == tokenNumeric && s.typ == tokenNumeric {
		fi, _ := strconv.Atoi(f.val)
		si, _ := strconv.Atoi(s.val)

		ni := fi / si

		return *tokenizeInt(ni), nil
	}
	return token{}, errors.New("cannot divid non-numeric tokens")
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func Split(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
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
