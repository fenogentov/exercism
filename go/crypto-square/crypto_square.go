package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(str string) string {
	if len(str) == 0 {
		return ""
	}
	str = strings.ToLower(str)
	f := func(r rune) rune {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			return r
		}
		return -1
	}
	str = strings.Map(f, str)
	c, r := rectangle(len(str))
	str += strings.Repeat(" ", (c*r - len(str)))
	var builder strings.Builder
	for i := 0; i < c; i++ {
		for in, s := range str {
			if in%c == i {
				builder.WriteString(string(s))
			}
		}
		if i != c-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

func rectangle(l int) (c, r int) {
	r = int(math.Sqrt(float64(l)))
	c = r
	if c*r == l {
		return c, r
	}
	c++
	if c*r < l {
		r++
	}
	return c, r
}
