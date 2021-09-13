// Package wordcount conducts frequency analysis of the proposal
package wordcount

import (
	"regexp"
	"strings"
	"unicode"
)

type Frequency map[string]int

// WordCount counts occurrences of each word in a phrase
func WordCountReg(txt string) Frequency {
	re := regexp.MustCompile("[\"\n\t\v\f\r\".,?!:`;&@$%^ ]+")
	txt = re.ReplaceAllString(txt, " ")
	re = regexp.MustCompile(" +'|' +|'$")
	txt = re.ReplaceAllString(txt, " ")
	txt = strings.Trim(txt, " ")
	s := strings.Split(txt, " ")

	t := make(Frequency)
	for _, w := range s {
		w = strings.ToLower(w)
		t[w]++
	}
	return t
}

func WordCount(txt string) Frequency {
	t := make(Frequency)
	f := func(c rune) bool {
		return c == ',' || c == '.' || unicode.IsSpace(c) || c == '$' || c == '^' || c == '!' || c == '%' ||
			c == '&' || c == '@' || c == ':'
	}
	fields := strings.FieldsFunc(txt, f)
	for _, w := range fields {
		w = strings.ToLower(strings.Trim(w, "'"))
		t[w]++
	}
	return t
}
