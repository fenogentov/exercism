package atbash

import (
	"strings"
	"unicode"
)

func Atbash(in string) string {
	in = strings.ToLower(in)
	inb := []rune(in)
	var cnt int
	var r strings.Builder
	for _, b := range inb {
		if unicode.IsDigit(b) {
			if cnt%5 == 0 && cnt != 0 {
				r.WriteString(" ")
			}
			r.WriteRune(b)
			cnt++
		} else if unicode.IsLetter(b) {
			b = 'z' - b%97
			if cnt%5 == 0 && cnt != 0 {
				r.WriteString(" ")
			}
			r.WriteRune(b)
			cnt++
		} else {
			continue
		}
	}
	return r.String()
}
