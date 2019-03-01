package reverse

// String revereses a string
func String(inputStr string) string {

	// Create a rune slice from the input string
	// Store the length of the runeSlice
	runeSlice := []rune(inputStr)
	runeSliceLength := len(runeSlice)

	// Swap the runeslice in place (Keep swapping the elements from each side with eachother)
	for i := 0; i < runeSliceLength/2; i++ {
		runeSlice[i], runeSlice[runeSliceLength-1-i] = runeSlice[runeSliceLength-1-i], runeSlice[i]
	}

	return string(runeSlice)
}
