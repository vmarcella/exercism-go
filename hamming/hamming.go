package hamming

import "errors"

// Distance calculates the distance between two strands of DNA
func Distance(a, b string) (int, error) {
	/*
	   Find the hamming distance between two strands of DNA

	   Arguments:
	       :a: strand of DNA of equal length to b
	       :b: strand of DNA of equal length to a
	*/
	if len(a) != len(b) {
		return 0, errors.New("DNA is not of equal length")
	}

	distance := 0

	// Compare each nucleotide of the DNA strands
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
