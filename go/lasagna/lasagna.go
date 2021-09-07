// Package lasagna is a solution to lerning #18 (exercism.io)
package lasagna

// OvenTime returns how many minutes the lasagna should be in the oven.
// According to the cooking book, the expected oven time in minutes is 40:
func OvenTime() int {
	return 40
}

// RemainingOvenTime returns how many minutes the lasagna still has to remain in the oven
func RemainingOvenTime(i int) int {
	return OvenTime() - i
}

// PreparationTime returns how many minutes you spent preparing the lasagna
func PreparationTime(l int) int {
	return l * 2
}

// ElapsedTime return how many minutes in total you've worked on cooking the lasagna
func ElapsedTime(l, t int) int {
	return l*2 + t
}
