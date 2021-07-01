// Package letter is a solution to lerning #10 (exercism.io)
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency // Frequency calculates the frequency of each rune of transmitted texts using concurrent computations
func ConcurrentFrequency(sl []string) FreqMap {

	m := FreqMap{}

	chFreqMap := make(chan FreqMap, 2)

	for _, s := range sl {

		go func(s string) {
			chFreqMap <- Frequency(s)
		}(s)
	}

	for range sl {
		c := <-chFreqMap
		for r, k := range c {
			m[r] += k

		}
	}

	return m
}
