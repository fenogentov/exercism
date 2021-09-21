package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(N int) bool {
	if N < 10 {
		return true
	}
	var summ int
	grade := len(strconv.Itoa(N))
	for i := 10; i < N*10; i *= 10 {
		a := N % i / (i / 10)
		summ += int(math.Pow(float64(a), float64(grade)))
	}
	return summ == N
}
