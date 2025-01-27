package solution1897

func makeEqual(words []string) bool {

	const (
		maxLetters = 26
	)

	freq := make([]int, maxLetters)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			freq[word[i]-'a']++
		}
	}

	for i := 0; i < maxLetters; i++ {
		if freq[i]%len(words) != 0 {
			return false
		}
	}

	return true
}
