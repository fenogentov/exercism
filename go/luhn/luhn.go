// Package luhn is a solution to lerning #7 (exercism.io)
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid validates the match to the Luhn algorithm
func Valid(in string) bool {

	in = strings.ReplaceAll(in, " ", "")

	if len(in) <= 1 {
		return false
	}

	summ := 0

	p := len(in)%2 == 0

	for _, s := range in {
		if !unicode.IsDigit(s) {
			return false
		}

		v, err := strconv.Atoi(string(s))
		if err != nil {
			return false
		}

		if p {
			v = 2 * v
			if v > 9 {
				v -= 9
			}
		}
		p = !p

		summ += v
	}

	return summ%10 == 0
}
