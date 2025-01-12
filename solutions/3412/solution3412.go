package solution3412

// ============================================================================
// 1920. Build Array from Permutation
// URL: https://leetcode.com/problems/build-array-from-permutation/
// ============================================================================

func calculateScore(s string) int64 {

	var score int64
	var freq [26][]int

	for i := 0; i < len(s); i++ {
		c1 := int(s[i] - 'a')
		c2 := int('z' - s[i])

		if len(freq[c2]) > 0 {
			score += int64(i - freq[c2][len(freq[c2])-1])
			freq[c2] = freq[c2][:len(freq[c2])-1]
		} else {
			freq[c1] = append(freq[c1], i)
		}
	}

	return score
}

func mirror(ch byte) byte {
	return 'z' - (ch - 'a')
}

// calculateScoreV1 -- this causes TLE (Time Limit Exceeded)
func calculateScoreV1(s string) int64 {

	var score int64

	isMirror := func(ch1 byte, ch2 byte) bool {

		c1 := ch1 - 'a'
		c2 := ch2 - 'a'

		if c1 < c2 && 25-c2 == c1 {
			return true
		}
		if c1 > c2 && 25-c1 == c2 {
			return true
		}

		return false
	}

	st := []byte(s)
	for i := 0; i < len(st); i++ {
		if st[i] == '-' {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			if st[j] == '-' {
				continue
			}
			if isMirror(st[i], st[j]) {
				st[i] = '-'
				st[j] = '-'
				score += int64(i - j)
			}
		}
	}

	return score
}
