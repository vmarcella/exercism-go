package wordcount

import "strings"

type Frequency map[string] int

func WordCount(sentence string) Frequency {
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
