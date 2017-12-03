// Package luhn determines if a number is a number is valid according to the Luhn formula
package luhn

import (
	"strings"
	"strconv"
)

// Valid checks the validity of a given string
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if len(s) > 1 {
		sum := 0
		for i,d := range(strings.Split(s,"")) {
			num, err := strconv.Atoi(d)
			if err != nil {
				return false
			}
			if i % 2 == len(s) % 2 {
				if 2*num  > 9 {
					sum += 2*num - 9
				} else {
					sum += 2*num
				}
			} else {
				sum += num
			}
		}
		if sum % 10 == 0 {
			return true
		}
	}
	return false
}
