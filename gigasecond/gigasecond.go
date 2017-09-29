// Package gigasecond implements a function to calculate the moment when someone has lived for 10^9 seconds
package gigasecond

// import path for the time package from the standard library
import (
	"time"
	"math"
)

const testVersion = 4

// AddGigasecond takes a time and returns the time 10^9 seconds later
func AddGigasecond(t0 time.Time) time.Time {
	return time.Unix(t0.Unix() + int64(math.Pow(10,9)), 0).UTC()
}
