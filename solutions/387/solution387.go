package solution387

// ============================================================================
// 387. First Unique Character in a String
// URL: https://leetcode.com/problems/first-unique-character-in-a-string/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_firstUniqChars-24                   	466539740	         6.450 ns/op	       0 B/op	       0 allocs/op
	Benchmark_firstUniqChars_first_approach-24    	 2505699	       478.0 ns/op	      48 B/op	       2 allocs/op
	PASS

*/

func firstUniqChar(s string) int {
	repeat := false
	for i := 0; i < len(s); i++ {
		repeat = false
		for j := 0; j < len(s); j++ {
			if i != j && s[i] == s[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			return i
		}
	}
	return -1
}

func firstUniqChar_first_approach(s string) int {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ch1 := byte(s[i])
		_, ok := m1[ch1]
		if !ok {
			m1[ch1] = 1
			m2[ch1] = i
		} else {
			m1[ch1]++
		}
		if m1[ch1] >= 2 {
			m2[ch1] = -1
		}
	}
	var mval = 1<<63 - 1
	var found = false
	for k, v := range m1 {
		if v == 1 && m2[k] != -1 {
			if m2[k] < mval {
				mval = m2[k]
				found = true
			}
		}
	}
	if found {
		return mval
	}
	return -1
}
