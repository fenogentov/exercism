package forth

import (
	"fmt"
	"strconv"
	"strings"
)

var operators map[string]interface{} = map[string]interface{}{"+": nil, "-": nil, "*": nil, "/": nil}

func Forth(inp []string) ([]int, error) {
	var buf []int
	for _, z := range inp {
		b := strings.Split(z, " ")
		for _, s := range b {
			fmt.Println(s)
			n, ok := strconv.Atoi(s)
			if ok == nil {
				buf = append(buf, n)
			}

			if _, ok := operators[s]; ok {
				oper(buf, s)
			}
			fmt.Println(buf)
		}
	}

	return []int{1, 2, 3, 4, 5}, nil

}

func oper(buf []int, op string) {
	l := len(buf)

	switch op {
	case "+":
		buf[l-2] = buf[l-2] + buf[l-1]
		buf = buf[:l-1]
	}
	fmt.Println(buf, l)
}
