package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// Our concurrent frequency counter function
func cFrequency(s string, channel chan FreqMap)  {
    m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	channel <- m
}

// Concurrent frequency go test -v --bench . --benchmem
func ConcurrentFrequency(strings []string) FreqMap {
    // Create the output channel and map
    var outputChannel = make(chan FreqMap)
    var outputMap = FreqMap{}

    // run the frequency count concurrently
    for _, currentString := range strings {
        go cFrequency(currentString, outputChannel)
    }
  
    // Iterate over the total list of strings
    for range strings {
        currentMap := <-outputChannel
        // Iterate over the map
        for letter, freq := range currentMap {
            outputMap[letter] += freq
        }
    }
    return outputMap 
}
