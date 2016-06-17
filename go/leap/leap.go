package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear checks years
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}
