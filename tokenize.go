package evaluator

import (
	"errors"
	"strconv"
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

type Token struct {
	typ tokenType
	val string
}

func TokenizeInt(i int) *Token {
	s := strconv.Itoa(i)
	return &Token{
		tokenNumeric,
		s,
	}
}

func TokenizeString(s string) (Token, error) {
	switch {
	case isNumeric(s):
		return Token{tokenNumeric, s}, nil
	case s == "+":
		return Token{tokenOperatorAdd, s}, nil
	case s == "-":
		return Token{tokenOperatorSub, s}, nil
	case s == "*":
		return Token{tokenOperatorMulti, s}, nil
	case s == "/":
		return Token{tokenOperatorDiv, s}, nil
	}
	return Token{tokenError, ""}, errors.New("unknown token")
}

func AddTwoTokens(f Token, s Token) (Token, error) {
	if f.typ == tokenNumeric && s.typ == tokenNumeric {
		fi, _ := strconv.Atoi(f.val)
		si, _ := strconv.Atoi(s.val)

		ni := fi + si

		return *TokenizeInt(ni), nil
	}
	return Token{}, errors.New("cannot add non-numeric tokens")
}
func SubTwoTokens(f Token, s Token) (Token, error) {
	if f.typ == tokenNumeric && s.typ == tokenNumeric {
		fi, _ := strconv.Atoi(f.val)
		si, _ := strconv.Atoi(s.val)

		ni := fi - si

		return *TokenizeInt(ni), nil
	}
	return Token{}, errors.New("cannot subtract non-numeric tokens")
}
func MultipleTwoTokens(f Token, s Token) (Token, error) {
	if f.typ == tokenNumeric && s.typ == tokenNumeric {
		fi, _ := strconv.Atoi(f.val)
		si, _ := strconv.Atoi(s.val)

		ni := fi * si

		return *TokenizeInt(ni), nil
	}
	return Token{}, errors.New("cannot multiple non-numeric tokens")
}
func DividTwoTokens(f Token, s Token) (Token, error) {
	if f.typ == tokenNumeric && s.typ == tokenNumeric {
		fi, _ := strconv.Atoi(f.val)
		si, _ := strconv.Atoi(s.val)

		ni := fi / si

		return *TokenizeInt(ni), nil
	}
	return Token{}, errors.New("cannot divid non-numeric tokens")
}
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
