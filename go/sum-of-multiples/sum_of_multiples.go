// Package luhn is a solution to lerning #7.1 (exercism.io)
package summultiples

// SumMultiples calculates the sum of all unique multiples from the slice
func SumMultiples(limit int, divisors ...int) (summ int) {
	summ = 0
	for l := 1; l < limit; l++ {
		for _, k := range divisors {
			if k == 0 {
				continue
			}
			if l%k == 0 {
				summ += l
				break
			}
		}
	}
	return
}
