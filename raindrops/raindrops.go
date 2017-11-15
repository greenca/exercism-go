// Package raindrops implements a solution to the exercism Raindrops exercise
package raindrops

import (
	"fmt"
	"math"
)

// Convert returns a string based on the factors of the integer argument
func Convert(i int)string {
	out := ""
	if math.Mod(float64(i), 3) == 0 {
		out += "Pling"
	}
	if math.Mod(float64(i), 5) == 0 {
		out += "Plang"
	}
	if math.Mod(float64(i), 7) == 0 {
		out += "Plong"
	}
	if out != "" {
		return out
	}
	return fmt.Sprintf("%d", i)
}
