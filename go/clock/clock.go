package clock

import "fmt"

const testVersion = 3

// Clock has minutes
type Clock struct {
	hour, minute int
}

// New Clock with hour and minute
func New(hour, minute int) Clock {
	c := Clock{hour, minute}
	return c
}

// Clock to string
func (c Clock) String() string {
	return fmt.Sprintf("%d:%d", c.hour, c.minute)
}

// Add minutes to Clock
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c
}
