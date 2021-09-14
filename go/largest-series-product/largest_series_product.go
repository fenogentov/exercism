package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(str string, q int) (int, error) {
	var num []int
	var max int
	if q > len(str) || q < 0 {
		return -1, errors.New("error span")
	}
	for _, s := range str {
		n, err := strconv.Atoi(string(s))
		if err != nil {
			return -1, err
		}
		num = append(num, n)
	}
	if q == len(str) {
		return Product(num, q), nil
	}
	for i := 0; i <= (len(num) - q); i++ {
		p := Product(num[i:i+q], q)
		if p > max {
			max = p
		}
	}
	return max, nil

}

func Product(i []int, q int) int {
	p := 1
	for j := 0; j < q; j++ {
		p *= i[j]
	}
	return p
}
