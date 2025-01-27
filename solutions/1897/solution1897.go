package solution1897

func makeEqual(words []string) bool {

	const (
		maxLetters = 26
	)

	freq := make([]byte, maxLetters)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			freq[byte(word[i])-'a']++
		}
	}

	for _, c := range freq {
		if int(c)%len(words) != 0 {
			return false
		}
	}

	return true
}
