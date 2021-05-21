// Package strand is a solution to lerning #3.1 (exercism.io)
package strand

import "strings"

// ToRNA builds up a DNA-based RNA strand
func ToRNA(dna string) string {
	var rna strings.Builder
	rna.Grow(len(dna))

	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'G':
			rna.WriteString("C")
		case 'C':
			rna.WriteString("G")
		case 'T':
			rna.WriteString("A")
		case 'A':
			rna.WriteString("U")
		}
	}

	return rna.String()
}
