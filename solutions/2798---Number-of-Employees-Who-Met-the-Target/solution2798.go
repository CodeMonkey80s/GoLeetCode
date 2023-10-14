package solution2798

// ============================================================================
// 2798. Number of Employees Who Met the Target
// URL: https://leetcode.com/problems/number-of-employees-who-met-the-target/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2798---Number-of-Employees-Who-Met-the-Target
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_numberOfEmployees-24    	1000000000	         0.8484 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	ans := 0
	for _, v := range hours {
		if v >= target {
			ans++
		}
	}
	return ans
}
