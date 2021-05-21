// Package accumulate is a solution to lerning #4.1 (exercism.io)
package accumulate

// Accumulate processes a slice of words according to the passed function
func Accumulate(str []string, fn func(string) string) []string {
	rez := []string{}
	for _, s := range str {
		rez = append(rez, fn(s))
	}
	return rez
}
