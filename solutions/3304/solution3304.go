package solution3304

// ============================================================================
// 3304. Find the K-th Character in String Game I
// URL: https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3304---Find-the-K-th-Character-in-String-Game-I
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_kthCharacter
	Benchmark_kthCharacter-24    	34771824	        38.08 ns/op	       8 B/op	       3 allocs/op
	PASS
*/

func kthCharacter(k int) byte {

	output := make([]byte, 0, 501)
	output = append(output, 'a')

	for i := 0; i < k; i++ {
		s := make([]byte, len(output))
		copy(s, output)

		for j := 0; j < len(s); j++ {
			s[j]++
		}

		output = append(output, s...)
		if len(output) > k {
			break
		}
	}

	return byte(output[k-1])
}
