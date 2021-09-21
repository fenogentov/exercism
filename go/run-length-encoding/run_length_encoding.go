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
	for i, r := range str {
		cnt++
		if len(str) == i+1 || str[i] != str[i+1] {
			if cnt > 1 {
				enc.WriteString(strconv.Itoa(cnt))
				enc.WriteRune(r)
			} else {
				enc.WriteRune(r)
			}
			cnt = 0
		}
	}
	return enc.String()
}

func RunLengthDecode(str string) string {
	var cntStr string
	var dec strings.Builder
	for i, r := range str {
		if len(str) > i+1 && unicode.IsDigit(r) {
			cntStr += string(r)
			continue
		}
		if cntStr == "" {
			dec.WriteRune(r)
		} else {
			cnt, err := strconv.Atoi(cntStr)
			if err != nil {
				log.Println(err)
			}
			dec.WriteString(strings.Repeat(string(r), cnt))
		}
		cntStr = ""
	}
	return dec.String()
}
