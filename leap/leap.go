// Package leap implements solution to exercism Leap exercise
package leap

const testVersion = 3

// IsLeapYear returns a boolean indicating whether the argument
// is a leap year
func IsLeapYear(year int) bool {
	return (year % 4 == 0) && (year % 100 != 0 || year % 400 == 0)
}

