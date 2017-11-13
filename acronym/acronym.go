// Package acronym implements a function to produce the acronym of a phrase
package acronym

import (
	"strings"
)

// Abbreviate takes an string and returns an upper-case acronym from the first letter of each word
func Abbreviate(s string) string {
	words := strings.Split(strings.Replace(s, "-", " ", -1), " ")
	out := ""
	for _, w := range words {
		out += string(w[0])
	}
	return strings.ToUpper(out)
}
