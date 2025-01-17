package solution1108

import "strings"

// ============================================================================
// 1108. Defanging an IP Address
// URL: https://leetcode.com/problems/defanging-an-ip-address/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1108
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_defang_stdlib
	Benchmark_defang_stdlib-24    	24414080	        50.09 ns/op	      16 B/op	       1 allocs/op
	Benchmark_defangV2
	Benchmark_defangV2-24    		21272322	        52.10 ns/op	      24 B/op	       2 allocs/op
	PASS
*/

func defangIPaddrV2(address string) string {
	var sb strings.Builder
	for _, ch := range []byte(address) {
		switch {
		case ch == '.':
			sb.WriteString("[.]")
		default:
			sb.WriteByte(ch)
		}
	}
	return sb.String()
}

func defangIPaddrV1(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
