package clock

import "fmt"

const testVersion = 4

// Clock has minutes
type Clock struct {
	hour, minute int
}

//GetRemainderQuotient is over MINUTES minutes or HOURS hours
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

// HOURS in a day
const HOURS = 24

// MINUTES in an hour
const MINUTES = 60

// New Clock with hour and minute
func New(hour, minute int) Clock {

	if (hour >= 0 && hour < HOURS) && (minute >= 0 && minute < MINUTES) {
		return Clock{hour, minute}
	}

	hr, min := hour, minute

	if minute < 0 {
		remMin, qntMin := GetRemainderQuotient(Abs(minute), MINUTES)
		min = MINUTES - remMin
		hr -= 24 - (qntMin + 1)
	}

	if hour < 0 {
		remHr, _ := GetRemainderQuotient(Abs(hour), HOURS)
		hr = HOURS - remHr
	}

	minRemainder, minQuotient := GetRemainderQuotient(min, MINUTES)

	if minQuotient > 0 {
		hr += minQuotient
		min = minRemainder
	}

	hrRemainder, hrQuotient := GetRemainderQuotient(hr, HOURS)

	if hrQuotient > 0 {
		hr = hrRemainder
	}

	if hr == HOURS {
		hr = 0
	}

	hr, min = Abs(hr), Abs(min)

	fmt.Printf("Passed: hour %d, minute %d \n", hour, minute)
	fmt.Printf("Calc: minRemainder %d, minQuotient %d \n", minRemainder, minQuotient)
	fmt.Printf("Calc: hrRemainder %d, hrQuotient %d \n", hrRemainder, hrQuotient)
	fmt.Printf("Return: hr %d, min %d \n", hr, min)

	return Clock{hr, min}
}

// Clock to string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add minutes to Clock
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes

	rem, qnt := GetRemainderQuotient(c.minute, MINUTES)

	if qnt > 0 {
		c.hour += qnt
		c.minute = rem
	}

	return c
}
