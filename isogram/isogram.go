package isogram

import "unicode"

// IsIsogram checks if a string is an isogram
func IsIsogram(inputStr string) bool {
	// Create a map as a pseudo set
	set := make(map[rune]bool)

	// Iterate over the input string
	for _, char := range inputStr {
		// map the rune to it's lowercase counterpart (if any)
		char := unicode.ToLower(char)

		// Check if the char is found
		if _, found := set[char]; found {

			// if the char is within the range of the alphabetical characters,
			// It has been seen before and thus the word is no longer an isogram
			if !(char < 'a' || char > 'z') {
				return false
			}
		}
		// Mark the char as seen
		set[char] = true
	}

	// Return true if we're able to iterate through the entire string without any hiccups
	return true
}
