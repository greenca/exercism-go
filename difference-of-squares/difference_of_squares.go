// Package diffsquares calculates the difference between the square of the sum of
// the first n natural numbers and the sum of the squares of the first n natural numbers
package diffsquares

// SquareOfSums returns the square of the sum of the first n natural numbers
func SquareOfSums(n int) int {
	sum := 0
	for i:=1; i<=n; i++ {
		sum += i
	}
	return sum*sum
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	sum := 0
	for i:=1; i<=n; i++ {
		sum += i*i
	}
	return sum
}

// Difference returns the difference between SquareOfSums and SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
