// Package leap is a solution to lerning #2.1 (exercism.io)
package leap

// IsLeapYear determines the sign of a leap year
func IsLeapYear(year int) bool {

	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}
