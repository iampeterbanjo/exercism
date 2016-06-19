package clock

import "fmt"

const testVersion = 4

// Clock has minutes
type Clock struct {
	hour, minute int
}

// New Clock with hour and minute
func New(hour, minute int) Clock {

	if (hour >= 0 && hour < MaxHours) && (minute >= 0 && minute < MaxMinutes) {
		return Clock{hour, minute}
	}

	min, additionalHours := GetMinutes(minute)
	hr := GetHours(additionalHours + hour)

	return Clock{hr, min}
}

// Clock to string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add minutes to Clock
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes

	return New(c.hour, c.minute)
}

// MaxHours in a day
const MaxHours = 24

// MaxMinutes in an hour
const MaxMinutes = 60

// CheckHours are between 0 and 24
func CheckHours(hours int) int {
	if hours == MaxHours {
		return 0
	}

	hr, _ := GetRemainderQuotient(hours, MaxHours)

	if hr < 0 {
		return MaxHours - Abs(hours)
	}

	return hr
}

// GetHours given minutes and hours
func GetHours(hours int) int {
	hr, _ := GetRemainderQuotient(hours, MaxHours)

	return CheckHours(hr)
}

// GetMinutes and return left over hours
func GetMinutes(minutes int) (min int, hr int) {
	min, hr = GetRemainderQuotient(minutes, MaxMinutes)

	if min < 0 {
		min = MaxMinutes - Abs(min)
		hr = hr - 1
	}

	return min, hr
}

//GetRemainderQuotient is over MaxMinutes minutes or MaxHours hours
func GetRemainderQuotient(num, divisor int) (remainder int, quotient int) {
	r := num % divisor
	q := int(num / divisor)

	return r, q
}

//Abs ..olute value for int
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
