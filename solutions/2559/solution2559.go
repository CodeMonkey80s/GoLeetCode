package solution2559

// ============================================================================
// 2559. Count Vowels Strings in Ranges
// URL: https://leetcode.com/problems/count-vowel-strings-in-ranges/
// ============================================================================

/*
	goarch: amd64
	pkg: GoLeetCode/solutions/2559
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_vowelString
	Benchmark_vowelString-24    	24686024	        50.81 ns/op	      72 B/op	       2 allocs/op
	PASS
*/

func vowelString(words []string, queries [][]int) []int {

	isVowel := func(s byte) bool {
		return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u'
	}

	var n int
	valid := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		if isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
			n++
		}
		valid[i] = n
	}

	var a int
	var b int
	var c int
	output := make([]int, 0, len(queries))
	for _, q := range queries {

		a = q[0] - 1
		b = q[1]
		if a < 0 {
			c = valid[b]
		} else {
			c = valid[b] - valid[a]
		}

		output = append(output, c)
	}

	return output
}
