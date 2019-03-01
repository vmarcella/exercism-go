package luhn

import "strconv"
import "strings"

func Valid(luhn string) bool {
	// remove all blank characters
	luhn = strings.Replace(luhn, " ", "", -1)

	// Check if the length of the lune is less than one
	if len(luhn) <= 1 {
		return false
	}

	// Keep track of which ints to double
	double := false

	// Keep track of the luhn sum
	sum := 0

	// Iterate over the luhn backwards to create the sum
	// Double every other number starting with the second number iterated over
	for index := len(luhn) - 1; index >= 0; index -= 1 {
		currentNum, err := strconv.Atoi(string(luhn[index]))
		if err == nil {
			if double {
				currentNum *= 2
				if currentNum > 9 {
					currentNum -= 9
				}
			}
			sum += currentNum
			double = !double
		} else {
			return false
		}
	}

	return sum%10 == 0
}
