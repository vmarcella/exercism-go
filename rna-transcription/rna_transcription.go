package strand

func ToRNA(dna string) string {
	// Create RNA mapping
	RNAComplement := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	// Create output RNA rune slice
	outputRNA := make([]rune, len(dna))

	// Map DNA to RNA
	for pos, char := range dna {
		outputRNA[pos] = RNAComplement[char]
	}

	// Return the string of the outputRNA rune slice
	return string(outputRNA)
}
