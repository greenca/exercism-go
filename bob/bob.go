// Package bob implements a solution to the exercism Bob exercise
package bob

import (
	"strings"
)

// Hey returns a string containing Bob's response to a given remark
func Hey(remark string) string {

	r := strings.TrimSpace(remark)

	if len(r) == 0 {
		return "Fine. Be that way!"
	}

	if r == strings.ToUpper(r) && r != strings.ToLower(r) {
		return "Whoa, chill out!"
	}

	if r[len(r)-1] == '?' {
		return "Sure."
	}
	
	return "Whatever."
}
