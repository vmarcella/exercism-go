package diffsquares

// SquareOfSum Calculate the square of the entire sum
func SquareOfSum(integer int) int {
	sum := 0
	for i := 1; i < integer+1; i += 1 {
		sum += i
	}

	return sum * sum
}

// SumOfSquares Calculate the sum of the squares
func SumOfSquares(integer int) int {
	sum := 0

	for i := 0; i < integer+1; i += 1 {
		sum += i * i
	}

	return sum
}

// Difference difference of the square of sums and the sum of the squares
func Difference(integer int) int {
	squareOfSum := SquareOfSum(integer)
	sumOfSquare := SumOfSquares(integer)

	return squareOfSum - sumOfSquare
}
