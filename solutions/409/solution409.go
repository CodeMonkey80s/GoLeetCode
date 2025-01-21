package solution409

func longestPalindrome(s string) int {

	if len(s) == 1 {
		return 1
	}

	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}

	var add int
	var longest int
	for _, c := range freq {
		switch {
		case c%2 == 0:
			longest += c
		case c%2 != 0:
			longest += c - 1
			add = 1
		}
	}

	return longest + add
}
