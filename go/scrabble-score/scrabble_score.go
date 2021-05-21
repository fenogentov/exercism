// Package scrabble is a solution to lerning #5 (exercism.io)
package scrabble

import "strings"

//Score calculating scrabble score by word
func Score(s string) int {
	s = strings.ToUpper(s)
	rez := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			rez += 1
		case 'D', 'G':
			rez += 2
		case 'B', 'C', 'M', 'P':
			rez += 3
		case 'F', 'H', 'V', 'W', 'Y':
			rez += 4
		case 'K':
			rez += 5
		case 'J', 'X':
			rez += 8
		case 'Q', 'Z':
			rez += 10
		default:
		}
	}
	return rez
}
