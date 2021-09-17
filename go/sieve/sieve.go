package sieve

import "math"

func Sieve(N int) (output []int) {

	m := make([]bool, N-1)
	for i := 0; i < N; i++ {
		m[N] = true
	}

	// Remove multiples of primes starting from 2, 3, 5,...
	for i := 2; i <= int(math.Sqrt(float64(N))); i++ {
		if m[i] {
			for j := i * i; j < N; j += i {
				m[j] = false
			}
		}
	}
	for i := 2; i < N; i++ {
		if m[i] {
			output = append(output, i)
		}
	}
	return output
}
