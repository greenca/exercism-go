// Package grains calculates the number of grains on wheat on a chessboard given that the number on each square doubles
package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains of wheat on a given square
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Invalid square number")
	}
	return uint64(math.Pow(2, float64(n-1))), nil
}

// Total returns the number of grains on wheat on the whole chessboard
func Total() uint64 {
	var sum uint64
	sum = 0
	for n:=1; n<=64; n++ {
		s, _ := Square(n)
		sum += s
	}
	return sum
}
