// Package isbn is ISBN verifier
package isbn

import (
	"errors"
	"strconv"
	"strings"
)

// IsValidISBN verifies ISBN-10
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	var r int
	for i, s := range isbn {

		if i == 9 && s == 'X' {
			r += 10
		} else {
			n, err := strconv.Atoi(string(isbn[i]))
			if err != nil {
				return false
			}
			r += n * (10 - i)
		}
	}
	return r%11 == 0
}

// ISBN10to13 convert ISBN-10 to ISBN-13
func ISBN10to13(isbn10 string) (isbn13 string, err error) {
	if !IsValidISBN(isbn10) {
		return "", errors.New("no valid ISBN10")
	}
	isbn := strings.ReplaceAll(isbn10, "-", "")
	isbn = "978" + isbn

	var c int
	var r int
	for i := 0; i < (len(isbn) - 1); i++ {
		c = 1
		n, err := strconv.Atoi(string(isbn[i]))
		if err != nil {
			return "", err
		}
		if (i+1)%2 == 0 {
			c = 3
		}
		r += n * c
	}
	r = 10 - (r % 10)
	if r == 10 {
		r = 0
	}
	crc := strconv.Itoa(r)
	isbn13 = "978-" + isbn10[:(len(isbn10)-1)] + crc
	return isbn13, nil
}

// IsValidISBN13 verifies ISBN-13
func IsValidISBN13(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	x13, err := strconv.Atoi(string(isbn[12]))
	if len(isbn) != 13 || err != nil {
		return false
	}

	var r int
	var c int
	for i := 0; i < (len(isbn) - 1); i++ {
		c = 1
		n, err := strconv.Atoi(string(isbn[i]))
		if err != nil {
			return false
		}
		if (i+1)%2 == 0 {
			c = 3
		}
		r += n * c
	}
	r = 10 - (r % 10)
	if r == 10 {
		r = 0
	}

	return r == x13
}
