// Package rotationalcipher implements a rotational cipher
package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher encryption using the librarys unicode && strings.Map
func RotationalCipher(str string, shiftKey int) string {
	f := func(r rune) rune {
		switch {
		case !unicode.IsLetter(r):
			return r
		case unicode.IsLower(r):
			return 'a' + (r-'a'+rune(shiftKey))%26

		case unicode.IsUpper(r):
			return 'A' + (r-'A'+rune(shiftKey))%26
		}
		return r
	}
	return strings.Map(f, str)
}

// RotationalCipher encryption using trings.Builder && brute force
func RotationalCipherV1(str string, shiftKey int) string {
	var res strings.Builder
	for i := range str {

		s := str[i]
		if s < 64 || (s > 90 && s < 97) || s > 122 {
			res.WriteByte(s)
			continue
		}
		b := s + byte(shiftKey)
		if s >= 64 && s <= 90 && b > 90 {
			b -= 26
		}
		if s >= 97 && s <= 122 && b > 122 {
			b -= 26
		}
		res.WriteByte(b)
	}
	return res.String()
}
