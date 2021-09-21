package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	errIncorrectNumber = errors.New("error incorrect number")
)

func Number(num string) (string, error) {
	f := func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}
	n := strings.Map(f, num)
	if len(n) < 10 || len(n) > 11 {
		return "", errIncorrectNumber
	}
	if len(n) == 11 {
		if n[0] != '1' {
			return "", errIncorrectNumber
		}
		n = n[1:]
	}
	if n[0] == '0' || n[0] == '1' || n[3] == '0' || n[3] == '1' {
		return "", errIncorrectNumber
	}
	return n, nil
}

func AreaCode(num string) (string, error) {
	n, err := Number(num)
	if err != nil {
		return "", err
	}
	return n[:3], nil

}

func Format(num string) (string, error) {
	n, err := Number(num)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", n[0:3], n[3:6], n[6:10]), nil
}
