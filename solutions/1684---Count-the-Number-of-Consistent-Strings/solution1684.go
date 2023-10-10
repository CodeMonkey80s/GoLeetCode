package solution1684

// ============================================================================
// 1684. Count the Number of Consistent Strings
// URL: https://leetcode.com/problems/count-the-number-of-consistent-strings/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1684---Count-the-Number-of-Consistent-Strings
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countConsistentStrings_map-24      	92496740	        65.10 ns/op	       0 B/op	       0 allocs/op
	Benchmark_countConsistentStrings_array-24    	390015631	        15.45 ns/op	       0 B/op	       0 allocs/op
	Benchmark_countConsistentStrings_slice-24    	345170995	        17.16 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func countConsistentStrings(allowed string, words []string) int {
	ans := 0
	m := make(map[rune]int, len(allowed))
	for _, char := range allowed {
		m[char] = 1
	}
outer:
	for _, word := range words {
		for _, char := range word {
			_, ok := m[char]
			if !ok {
				continue outer
			}
		}
		ans++
	}
	return ans
}

func countConsistentStrings_array(allowed string, words []string) int {
	ans := 0
	m := [26]int{}
	for _, char := range allowed {
		m[char-'a'] = 1
	}
outer:
	for _, word := range words {
		for i := range word {
			if m[word[i]-'a'] == 0 {
				continue outer
			}
		}
		ans++
	}
	return ans
}

func countConsistentStrings_slice(allowed string, words []string) int {
	ans := 0
	m := make([]int, 26)
	for _, char := range allowed {
		m[char-'a'] = 1
	}
outer:
	for _, word := range words {
		for i := range word {
			if m[word[i]-'a'] == 0 {
				continue outer
			}
		}
		ans++
	}
	return ans
}
