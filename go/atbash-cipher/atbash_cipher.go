package atbash

import (
	"strings"
)

func Atbash(in string) string {
	in = strings.ToLower(in)
	inb := []byte(in)
	var cnt int
	var r strings.Builder
	for _, b := range inb {
		if b > 47 && b < 58 {
			if cnt%5 == 0 && cnt != 0 {
				r.WriteString(" ")
			}
			r.WriteByte(b)
			cnt++
		} else if b > 96 && b < 123 {
			b = 122 - b%97
			if cnt%5 == 0 && cnt != 0 {
				r.WriteString(" ")
			}
			r.WriteByte(b)
			cnt++
		} else {
			continue
		}
	}
	return r.String()
}
