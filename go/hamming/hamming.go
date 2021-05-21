// Package hamming is a solution to lerning #3 (exercism.io)
package hamming

import (
	"errors"
)

var ErrDiffLenghtDNA = errors.New("error DNA sequences of different lengths")

// Distance calculates the Hamming distance between two strands of DNA.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, ErrDiffLenghtDNA
	}

	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
