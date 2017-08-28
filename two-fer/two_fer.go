// Package twofer implements solution to exercism Two Fer exercise
package twofer

import "fmt"

// ShareWith returns a string of the form "One for X, one for me."
// where X is provided by the argument. If the argument is an empty
// string, X is replaced by "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
