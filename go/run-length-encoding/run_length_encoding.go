package encode

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(str string) string {
	var cnt int
	var enc string
	for i := range str {
		cnt++
		if len(str) == i+1 || str[i] != str[i+1] {
			if cnt > 1 {
				enc += strconv.Itoa(cnt) + string(str[i])
			} else {
				enc += string(str[i])
			}
			cnt = 0
		}
	}
	return enc
}

func RunLengthDecode(str string) string {
	var cntStr string
	var dec string
	for i, s := range str {
		if len(str) > i+1 && unicode.IsDigit(rune(s)) {
			cntStr += string(s)
			continue
		}
		if cntStr == "" {
			dec += string(s)
		} else {
			cnt, err := strconv.Atoi(cntStr)
			if err != nil {
				log.Println(err)
			}
			dec += strings.Repeat(string(s), cnt)
		}
		cntStr = ""
	}
	return dec

}
