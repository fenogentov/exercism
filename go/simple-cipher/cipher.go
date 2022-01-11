package cipher

import (
	"regexp"

	"strings"
)

type Cipher interface {
	Encode(string) string

	Decode(string) string
}

type VigenereCipher string

var simplify = regexp.MustCompile(`[^a-z]`)

func NewCaesar() *VigenereCipher {
	return NewShift(3)
}

func NewShift(shift int) *VigenereCipher {
	if shift < 0 {
		shift += 26
	}
	return NewVigenere(string(rune('a' + shift)))
}

func NewVigenere(key string) *VigenereCipher {
	hasNonA := false
	for _, k := range key {
		if k < 'a' || k > 'z' {
			return nil
		}
		if k != 'a' {
			hasNonA = true
		}
	}
	if !hasNonA {
		return nil
	}
	cipher := VigenereCipher(key)
	return &cipher
}

func (c VigenereCipher) Encode(plain string) (encoded string) {
	simple := simplify.ReplaceAllString(strings.ToLower(plain), "")
	for i, r := range simple {
		newval := r + (rune(c[i%len(c)]) - 'a')
		if newval > 'z' {
			newval -= 26
		}
		encoded += string(newval)
	}
	return encoded
}

func (c VigenereCipher) Decode(encoded string) (plain string) {
	for i, r := range encoded {
		newval := r - (rune(c[i%len(c)]) - 'a')
		if newval < 'a' {
			newval += 26
		}
		plain += string(newval)
	}
	return plain
}
