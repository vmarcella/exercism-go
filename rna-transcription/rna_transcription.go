package strand

func ToRNA(dna string) string {
    RNAComplement := map[rune]rune{
        'G':'C',
        'C':'G',
        'T':'A',
        'A':'U',
    }

    outputRNA := make([]rune, len(dna))
    for pos, char := range dna {
        outputRNA[pos] = RNAComplement[char]
    }
    
    return string(outputRNA)
}
