// Package romannumerals provides tools for working with Roman numerals.
package romannumerals

import (
	"errors"
	"strings"
)

var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var numerals = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

// ToRomanNumeral function of converting Arabic numbers to Roman numerals.
func ToRomanNumeral(number int) (string, error) {
	var res strings.Builder
	if number < 1 || number > 3000 {
		return "", errors.New("value must be in the range 1 ... 3000")
	}
	for i := 0; i < 13; i++ {
		for number >= values[i] {
			number -= values[i]
			res.WriteString(numerals[i])
		}
	}
	return res.String(), nil
}
