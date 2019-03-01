package grains

import "errors"

// Square Return the amount of grains you'd get from a single squaree
func Square(square int) (uint64, error) {
	if square <= 0 {
		return 0, errors.New("you didn't provide a int greater than 0")
	} else if square > 64 {
		return 0, errors.New("you provided a int greater than 64")
	}

	grains := 1
	for i := 1; i < square; i++ {
		grains += grains
	}

	return uint64(grains), nil
}

// Total sums all the grains that would be on each square
func Total() uint64 {
	var total uint64
	// Iterate through the entire board
	for i := 1; i < 65; i++ {
		grains, _ := Square(i)
		total += uint64(grains)
	}

	return total
}
