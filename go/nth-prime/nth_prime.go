package prime

import "math"

func isPrime(p int) bool {
	if p%2 == 0 {
		return p == 2
	}
	m := int(math.Sqrt(float64(p))) + 1
	for t := 3; t < m; t += 2 {
		if p%t == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	for p := 2; ; p++ {
		if isPrime(p) {
			n--
			if n == 0 {
				return p, true
			}
		}
	}
}
