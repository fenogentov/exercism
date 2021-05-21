// Package raindrops is a solution to lerning #4 (exercism.io)
package raindrops

import (
	"strconv"
	"strings"
)

// Convert transforms the number to rain sounds
func Convert(i int) string {
	var o strings.Builder
	var flag bool

	if i%3 == 0 {
		o.WriteString("Pling")
		flag = true
	}
	if i%5 == 0 {
		o.WriteString("Plang")
		flag = true
	}
	if i%7 == 0 {
		o.WriteString("Plong")
		flag = true
	}
	if !flag {
		o.WriteString(strconv.Itoa(i))
	}

	return o.String()
}
