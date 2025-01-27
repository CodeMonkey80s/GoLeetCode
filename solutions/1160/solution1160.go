package solution1160

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1160
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countCharacters
	Benchmark_countCharacters-24    	39549944	        53.26 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countCharacters(words []string, chars string) int {

	const (
		maxLetters = 26
	)

	getFreqForWord := func(word string) []byte {
		freq := make([]byte, maxLetters)
		for i := 0; i < len(word); i++ {
			ch := word[i] - 'a'
			freq[ch]++
		}

		return freq
	}

	freqChars := getFreqForWord(chars)

	var sum int
loop:
	for _, word := range words {
		freqWord := getFreqForWord(word)
		for i := 0; i < maxLetters; i++ {
			if freqWord[i] > freqChars[i] {
				continue loop
			}
		}
		sum += len(word)
	}

	return sum
}
