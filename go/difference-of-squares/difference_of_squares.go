// Package diffsquares is a solution to lerning #7 (exercism.io)
package diffsquares

// SquareOfSum calculates the Square of the sum of n natural numbers
func SquareOfSum(n int) int {
	sum := (1 + n) * n / 2
	return sum * sum
}

//SumOfSquares calculates sum of the squares of the first n natural numbers is
func SumOfSquares(n int) int {
	squar := n * (n + 1) * (2*n + 1) / 6
	return squar
}

// Difference  between the square of the sum of the first n natural numbers and the sum of the squares of the first n natural numbers
func Difference(n int) int {
	return (SquareOfSum(n) - SumOfSquares(n))
}
