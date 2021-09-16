package encode

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(str string) string {
	var cnt int
	var enc strings.Builder
	for i, s := range str {
		cnt++
		if len(str) == i+1 || str[i] != str[i+1] {
			if cnt > 1 {
				enc.WriteString(strconv.Itoa(cnt))
				enc.WriteRune(s)
			} else {
				enc.WriteRune(s)
			}
			cnt = 0
		}
	}
	return enc.String()
}

func RunLengthDecode(str string) string {
	var cntStr string
	var dec strings.Builder
	for i, s := range str {
		if len(str) > i+1 && unicode.IsDigit(rune(s)) {
			cntStr += string(s)
			continue
		}
		if cntStr == "" {
			dec.WriteRune(s)
		} else {
			cnt, err := strconv.Atoi(cntStr)
			if err != nil {
				log.Println(err)
			}
			dec.WriteString(strings.Repeat(string(s), cnt))
		}
		cntStr = ""
	}
	return dec.String()
}
