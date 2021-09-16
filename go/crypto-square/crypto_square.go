package cryptosquare

import (
	"fmt"
	"strings"
	"unicode"
)

func Encode(str string) string {
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
	fmt.Println(str)
	var res strings.Builder

	for i := 0; i <= r; i++ {
		if i != 0 {
			res.WriteString(" ")
		}
		for j := 0; j < len(str); j += c {
			fmt.Println(i, j, i+j, len(str))
			res.WriteByte(str[i+j])
		}
		fmt.Println(res.String())
	}
	return res.String()
}

func rectangle(l int) (c, r int) {
	for r := 0; r < l; r++ {
		for c := r; c <= (r + 1); c++ {
			if c*r >= l {
				return c, r
			}
		}
	}
	return 0, 0
}
