// Package protein convert RNA to protein
package protein

import (
	"errors"
)

var ErrStop = errors.New("codon STOP")
var ErrInvalidBase = errors.New("invalid base codon")

// FromCodon translate codon sequences into protein
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA translate RNA sequences into proteins
func FromRNA(rna string) ([]string, error) {
	r := []string{}
	for i := 0; i < len(rna); i += 3 {
		codon, err := FromCodon(rna[i : i+3])
		if err == ErrInvalidBase {
			return r, err
		}
		if err == ErrStop {
			return r, nil
		}
		r = append(r, codon)
	}
	return r, nil
}
