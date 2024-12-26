package solution3271

// ============================================================================
// 3271. Hash Divided String
// URL: https://leetcode.com/problems/hash-divided-string/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3271
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_stringHashV1
Benchmark_stringHashV1-24    	45960301	        28.45 ns/op	       8 B/op	       1 allocs/op
	Benchmark_stringHashV2
	Benchmark_stringHashV2-24    	50958140	        22.71 ns/op	       8 B/op	       1 allocs/op
	PASS
*/

import "strings"

func stringHashV2(s string, k int) string {
	if k <= 0 || len(s)%k != 0 {
		return ""
	}

	var sb strings.Builder
	step := len(s) / k
	sb.Grow(step)

	for n := 0; n < step; n++ {
		sum := 0
		for i := 0; i < k; i++ {
			idx := n*k + i
			charValue := int(s[idx] - 'a')
			sum += charValue
		}
		hashedChar := rune((sum % 26) + 'a')
		sb.WriteRune(hashedChar)
	}

	return sb.String()
}

func stringHashV1(s string, k int) string {
	var sb strings.Builder
	step := len(s) / k
	for n := 0; n < step; n++ {
		sum := 0
		for i := 0; i < k; i++ {
			idx := n*k + i
			n := int(s[idx] - 'a')
			sum += n
		}
		hashedChar := sum % 26
		sb.WriteRune(rune(hashedChar + 'a'))
	}

	return sb.String()
}
