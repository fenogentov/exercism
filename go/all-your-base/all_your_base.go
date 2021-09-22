// Package allyourbase implements tools for converting a sequence of numbers
// from one base to any other base.
package allyourbase

import (
	"errors"
)

var (
	errInBase  = errors.New("input base must be >= 2")
	errOutBase = errors.New("output base must be >= 2")
	errDigit   = errors.New("all digits must satisfy 0 <= d < input base")
)

// ConvertToBase converts a sequence of numbers with base 'inBase'
// to a sequence of numbers with base 'outBase'
func ConvertToBase(inBase int, digits []int, outBase int) (out []int, err error) {
	if inBase < 2 {
		return nil, errInBase
	}
	if outBase < 2 {
		return nil, errOutBase
	}
	N, err := toTenBase(digits, inBase)
	if err != nil {
		return nil, err
	}
	out = toAnyBase(N, outBase)
	return out, nil
}

// intPow performs exponentiation for integers.
func intPow(digit, grade int) (res int) {
	if grade == 0 {
		return 1
	}
	if grade == 1 {
		return digit
	}
	res = digit
	for i := 1; i < grade; i++ {
		res *= digit
	}
	return res
}

// toTenBase converts a sequence of numbers with an arbitrary base gets a number with base 10.
func toTenBase(digits []int, inBase int) (N int, err error) {
	for i := 0; i < len(digits); i++ {
		if digits[i] < 0 || digits[i] > (inBase-1) {
			return 0, errDigit
		}
		N += digits[i] * intPow(inBase, (len(digits)-i-1))
	}
	return N, nil
}

// toAnyBase gets from base 10 a sequence of numbers with any base
func toAnyBase(digit, outBase int) (out []int) {
	if digit < outBase {
		out = append(out, digit)
		return
	}
	r := toAnyBase(digit/outBase, outBase)
	return append(r, digit%outBase)

}
