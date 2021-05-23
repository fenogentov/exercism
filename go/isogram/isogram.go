// Package isogram is a solution to lerning #6 (exercism.io)
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram defines is an isogram
func IsIsogram(s string) bool {
	s = strings.ToLower(s)

	mr := make(map[rune]bool)

	for _, r := range s {

		_, ok := mr[r]
		if ok && unicode.IsLetter(r) {
			return false
		}

		mr[r] = true
	}

	return true
}
