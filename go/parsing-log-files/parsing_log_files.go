package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	rgxp, _ := regexp.Compile(`(^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]).+`)
	return rgxp.MatchString(text)
}

func SplitLogLine(text string) []string {
	rgxp, _ := regexp.Compile(`\<\W*[^a-z]|[^a-z]\>`)
	return rgxp.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	rgxp, _ := regexp.Compile(`(?i)\".*password.*\"`)
	cnt := 0
	for _, s := range lines {
		if rgxp.MatchString(s) {
			cnt++
		}
	}
	return cnt
}

func RemoveEndOfLineText(text string) string {
	rgxp, _ := regexp.Compile(`end-of-line(\d+)`)
	return rgxp.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	rgxp := regexp.MustCompile(`(User\s+)(\w+)`)

	var finalResult []string

	for _, v := range lines {
		matches := rgxp.FindStringSubmatch(v)
		var nameToPrefix string

		if len(matches) > 0 {
			nameToPrefix = fmt.Sprintf("[USR] %s %s", matches[2], v)
		} else {
			nameToPrefix = v
		}

		finalResult = append(finalResult, nameToPrefix)
	}

	return finalResult
}
