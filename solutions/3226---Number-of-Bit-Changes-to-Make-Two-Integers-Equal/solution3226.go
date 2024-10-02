package solution3226

// ============================================================================
// 3226. Number of Bit Changes to Make Two Integers Equal
// URL: https://leetcode.com/problems/number-of-bit-changes-to-make-two-integers-equal/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3226---Number-of-Bit-Changes-to-Make-Two-Integers-Equal
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_minChanges
	Benchmark_minChanges-24    	266556715	         4.485 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func minChanges(n int, k int) int {

	if n == k {
		return 0
	}

	var ans int

	v := n
	for i := 0; i <= 32; i++ {

		val1 := v & (1 << i)
		val2 := k & (1 << i)

		if val1 > 0 && val2 == 0 {
			ans++
			v &^= 1 << i
			if v == k {
				return ans
			}
		}
	}

	return -1
}
