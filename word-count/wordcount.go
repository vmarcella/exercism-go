package wordcount

import "strings"
import "fmt"

// Frequency maps strings to integers
type Frequency map[string]int

// WordCount gets the word count of a given string
func WordCount(sentence string) Frequency {
	fmt.Println(sentence)
	sentence = strings.Replace(sentence, ",", " ", -1)

	wordList := strings.Split(sentence, " ")
	wordCount := Frequency{}

	for _, element := range wordList {
		_, found := wordCount[element]
		if found {
			wordCount[element]++
		} else {
			wordCount[element] = 1
		}
	}
	return wordCount
}
