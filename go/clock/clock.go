package clock

import "fmt"

const testVersion = 4

// Clock has minutes
type Clock struct {
	hour, minute int
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

// MaxHours in a day
const MaxHours = 24

// MaxMinutes in an hour
const MaxMinutes = 60

// New Clock with hour and minute
func New(hour, minute int) Clock {

	if (hour >= 0 && hour < MaxHours) && (minute >= 0 && minute < MaxMinutes) {
		return Clock{hour, minute}
	}

	totalMinutes := hour*MaxMinutes + minute
	min, mq := GetRemainderQuotient(totalMinutes, MaxMinutes)
	_, hq := GetRemainderQuotient(totalMinutes, MaxHours)
	hq += mq

	hr, _ := GetRemainderQuotient(hq, MaxHours)

	return Clock{hr, min}
}

// Clock to string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add minutes to Clock
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes

	rem, qnt := GetRemainderQuotient(c.minute, MaxMinutes)

	if qnt > 0 {
		c.hour += qnt
		c.minute = rem
	}

	return c
}
