package solution2423

// ============================================================================
// 2423. Remove Letter To Equalize Frequency
// URL: https://leetcode.com/problems/remove-letter-to-equalize-frequency/
// ============================================================================

// ! Nonfunctional // Does not work !

func equalFrequency(word string) bool {
	const m = 26
	freq := make([]int, m)
	for _, char := range word {
		idx := int(char - 'a')
		//fmt.Printf("idx = %d, %c\n", idx, char)
		freq[idx]++
	}
	isSame := func(freq []int, m int) bool {
		s := 0
		i := 0
		for i = 0; i < m; i++ {
			if freq[i] > 0 {
				s = freq[i]
				break
			}
		}
		for j := i + 1; j < m; j++ {
			if freq[j] > 0 && freq[j] != s {
				return false
			}
		}
		return true
	}
	if isSame(freq, m) {
		return true
	}
	for i := 0; i < m; i++ {
		//fmt.Printf("i=%d, %d\n", i, freq[i])
		if freq[i] > 0 {
			freq[i]--
			if isSame(freq, m) {
				return true
			}
			freq[i]++
		}
	}
	return false
}
