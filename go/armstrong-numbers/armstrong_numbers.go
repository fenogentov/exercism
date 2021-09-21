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
	grade := int(math.Log10(float64(N))) + 1
	for i := 10; i < N*10; i *= 10 {
		d := N % i / (i / 10)
		summ += intPow(d, grade)
	}
	return summ == N
}
func IsNumberV1(N int) bool {
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

func intPow(n, grade int) (pow int) {
	pow = n
	for i := 1; i < grade; i++ {
		pow *= n
	}
	return pow
}
