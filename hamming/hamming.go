// Package hamming implements a function to calculate the hamming distance between two strings
package hamming

import "fmt"

const testVersion = 6

// Distance takes two strings and returns the number of characters that differ between them
func Distance(a, b string) (int, error) {
	d := 0

	if len(a) != len(b) {
		return d, fmt.Errorf("String are different lengths")
	}

	for i:=0; i<len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}
