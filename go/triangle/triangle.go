// Package triangle is a solution to lerning #2.2 (exercism.io)
package triangle

import (
	"math"
)

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
	Deg             // degenerate
)

// KindFromSides returns the type of the triangle
func KindFromSides(a, b, c float64) Kind {

	var k Kind

	if (a <= 0 || b <= 0 || c <= 0) ||
		(a+b < c || a+c < b || c+b < a) ||
		(math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)) ||
		(math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1)) ||
		(math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1)) {
		return NaT
	}

	switch {
	case (a+b == c || a+c == b || b+c == a):
		k = Deg
	case (a == b) && (b == c):
		k = Equ
	case (a == b) || (a == c) || (b == c):
		k = Iso
	case (a != b) && (a != c) && (b != c):
		k = Sca
	}

	return k
}
