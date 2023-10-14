package solution1670

// ============================================================================
// 1672. Richest Customer Wealth
// URL: https://leetcode.com/problems/richest-customer-wealth/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1670---Richest-Customer-Wealth
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maximumWealth-24    	706966531	         2.525 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func maximumWealth(accounts [][]int) int {
	ans := 0
	for _, money := range accounts {
		sum := 0
		for i := 0; i < len(money); i++ {
			sum += money[i]
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
