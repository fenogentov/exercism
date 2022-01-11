package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.
// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	number int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.number)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction {
		amount *= 2
		err = nil
	}

	switch {
	case err != nil:
		return 0, err

	case amount < 0:
		return 0, errors.New("negative fodder")

	case cows == 0:
		return 0, errors.New("division by zero")

	case cows < 0:
		return 0, &SillyNephewError{number: int(cows)}

	default:
		return amount / float64(cows), err
	}
}
