package savings

// FixedInterestRate has a value of 5% (5/100)

// GetFixedInterestRate returns the FixedInterestRate constant
func GetFixedInterestRate() float32 {
	return 0.05
}

// DaysPerYear has a value of 365

// GetDaysPerYear returns the DaysPerYear constant
func GetDaysPerYear() int {
	return 365
}

// Jan-Dec have values of 1-12
const (
	_ = iota
	Jan
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

// GetMonth returns the value for the given month
func GetMonth(month int) int {
	return month
}

// AccNo type for a string - this is a stub type included to demonstrate how the untyped string constant can be used where this type is required
type AccNo string

var accno AccNo = "XF348IJ"

// AccountNo has a value of XF348IJ

// GetAccountNumber returns the AccountNo constant
func GetAccountNumber() AccNo {
	return accno
}
