// Package pangram determines if a sentence is a pangram
package pangram

import (
	"bytes"
)

func IsPangram(s string) bool {
	bs := []byte(s)
	for b := 65; b <= 90; b++ {
		ok := bytes.Contains(bs, []byte{byte(b)}) || bytes.Contains(bs, []byte{byte(b + 32)})
		if !ok {
			return false
		}
	}
	return true
}
