package solution3210

import "strings"

// ============================================================================
// 3210. Find the Encrypted String
// URL: https://leetcode.com/problems/find-the-encrypted-string/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3210---Find-the-Encrypted-String
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkGetEncryptedString
	BenchmarkGetEncryptedString-24    	52121913	        28.52 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func getEncryptedString(s string, k int) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for i := 0; i < len(s); i++ {
		sb.WriteRune(rune(s[(i+k)%len(s)]))
	}

	return sb.String()
}
