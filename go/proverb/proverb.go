// Package proverb is a solution to lerning #2.3 (exercism.io)
package proverb

// Proverb from a set of words
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	var proverb []string

	for i := 1; i < len(rhyme); i++ {
		proverb = append(proverb, "For want of a "+rhyme[i-1]+" the "+rhyme[i]+" was lost.")
	}

	proverb = append(proverb, "And all for the want of a "+rhyme[0]+".")

	return proverb
}
