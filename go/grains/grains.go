// Package grains is a solution to lerning #8 (exercism.io)
package grains

import "errors"

// Square raise the number 2 to the power
func Square(in int) (uint64, error) {
	if in < 1 || in > 64 {
		return 0, errors.New("invalid input value")
	}
	return 1 << (in - 1), nil
}

//Total calculation of the total amount of grains
func Total() uint64 {
	return 1<<64 - 1
}
