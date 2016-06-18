package clock

import (
	"fmt"
	"time"
)

const testVersion = 4

// Clock has minutes
type Clock struct {
	hour, minute int
}

// New Clock with hour and minute
func New(hour, minute int) Clock {
	fmt.Printf("Passed: hour %d, minute %d \n", hour, minute)
	t := time.Time{}
	a := t.Add(time.Duration(hour) * time.Hour)
	b := a.Add(time.Duration(minute) * time.Hour)

	c := Clock{hour, minute}
	c.hour, c.minute, _ = b.Clock()
	fmt.Printf("Return: c.hour %d, c.minute %d \n", c.hour, c.minute)

	return c
}

// Clock to string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add minutes to Clock
func (c Clock) Add(minutes int) Clock {
	t := time.Time{}
	x := t.Add(time.Duration(c.hour) * time.Hour)
	y := x.Add(time.Duration(c.minute) * time.Minute)
	z := y.Add(time.Duration(minutes) * time.Minute)

	c.hour, c.minute, _ = z.Clock()
	return c
}
