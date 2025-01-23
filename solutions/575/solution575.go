package solution575

func distributeCandies(candyType []int) int {

	n := len(candyType)
	types := n / 2

	freq := make(map[int]int, len(candyType))
	for i := 0; i < n; i++ {
		freq[candyType[i]]++
	}

	if types < len(freq) {
		return types
	}

	return len(freq)
}
