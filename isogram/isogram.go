// Package isogram implements a function to determine if a word or phrase is an isogram
package isogram

import "strings"

// IsIsogram takes a string and returns a boolean indicating if it is an isogram
func IsIsogram(str string) bool {
	s := strings.ToLower(str)
	for i, c := range(strings.Split(s, "")) {
		if c != " " && c != "-" && strings.Contains(s[i+1:], c) {
			return false
		}
	}
	return true
}
