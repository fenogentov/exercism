// Package variablelengthquantity implements Variable Length Quantity algorithm.
package variablelengthquantity

import (
	"errors"
)

// DecodeVarint decodes slice bytes to unsigned numbers.
func DecodeVarint(input []byte) (output []uint32, err error) {
	var n uint32
	for i, b := range input {
		a := b >> 7
		n = n<<7 + uint32(b&0x7f)
		if a == 1 && (i == len(input)-1) {
			return nil, errors.New("err")
		}
		if a == 0 {
			output = append(output, n)
			n = 0
		}
	}
	return output, nil
}

// EncodeVarint encodes a slice of numbers into a sequence of bytes.
func EncodeVarint(input []uint32) (output []byte) {
	for _, in := range input {
		if in == 0 {
			output = append(output, 0x00)
			continue
		}
		var buf []byte
		for in > 0 {
			b := byte(in & 0x7f)
			in >>= 7
			if len(buf) > 0 {
				b |= 0x80
			}
			buf = append(buf, b)
		}

		for i := len(buf); i > 0; i-- {
			output = append(output, buf[i-1])
		}
	}
	return output
}
