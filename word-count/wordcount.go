package wordcount

import "strings"
import "fmt"

type Frequency map[string]int

func WordCount(sentence string) Frequency {
	fmt.Println(sentence)
	sentence = strings.Replace(sentence, ",", " ", -1)
	word_list := strings.Split(sentence, " ")
	word_count := Frequency{}

	for _, element := range word_list {
		_, found := word_count[element]
		if found {
			word_count[element] += 1
		} else {
			word_count[element] = 1
		}
	}
	return word_count
}
