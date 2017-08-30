// Package clock implements a clock that handles time without dates
package clock

import "fmt"

const testVersion = 4

// Clock is a struct containing the time
type Clock struct {
	Hour int
	Minute int
}

// New returns a new Clock struct
func New(hour, minute int) Clock {
	hour += minute/60
	minute = minute % 60
	if minute < 0 {
		hour -= 1
		minute += 60
	}
	hour = hour % 24
	if hour < 0 {
		hour += 24
	}
	return Clock{
		Hour: hour,
		Minute: minute,
	}
}

// String returns a string representation of a Clock's time
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

// Add returns a new Clock with the specified number of minutes
// (positive or negative) added to the original Clock's time
func (c Clock) Add(minutes int) Clock {
	return New(c.Hour, c.Minute + minutes)
}
