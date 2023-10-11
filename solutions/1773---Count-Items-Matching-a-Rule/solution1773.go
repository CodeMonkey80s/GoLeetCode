package solution1773

// ============================================================================
// 1773. Count Items Matching a Rule
// URL: https://leetcode.com/problems/count-items-matching-a-rule/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1773---Count-Items-Matching-a-Rule
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countMatches-24    	347675583	         3.431 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ans := 0
	idx := 0
	for _, item := range items {
		switch ruleKey {
		case "type":
			idx = 0
		case "color":
			idx = 1
		case "name":
			idx = 2
		}
		if item[idx] == ruleValue {
			ans++
		}
	}

	return ans
}
