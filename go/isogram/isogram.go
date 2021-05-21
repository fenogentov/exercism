// Package isogram is a solution to lerning #6 (exercism.io)
package isogram

import (
	"strings"
)

// IsIsogram defines is an isogram
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	ch := make(map[byte]bool)
	for i := 0; i < len(s); i++ {

		_, ok := ch[s[i]]

		if ok && !(s[i] == ' ' || s[i] == '-') {
			return false
		}

		ch[s[i]] = true
	}

	return true
}
