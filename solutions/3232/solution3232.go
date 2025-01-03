package solution3232

// ============================================================================
// 3232. Find if Digit Game Can Be Won
// URL: https://leetcode.com/problems/find-if-digit-game-can-be-won/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3232---Find-if-Digit-Game-Can-Be-Won
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_canAliceWingV2
	Benchmark_canAliceWingV2-24    	1000000000	         2.556 ns/op	       0 B/op	       0 allocs/op
	Benchmark_canAliceWingV1
	Benchmark_canAliceWingV1-24    	279003758	         4.297 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func canAliceWinV2(nums []int) bool {

	var as, ad int

	for _, v := range nums {
		switch {
		case v >= 1 && v <= 9:
			as += v
		default:
			ad += v
		}
	}

	return as-ad != 0
}

func canAliceWinV1(nums []int) bool {

	var as, ad, bs, bd int

	for _, v := range nums {

		if v >= 1 && v <= 9 {
			as += v
		} else {
			bd += v
		}

		if v >= 10 && v <= 99 {
			ad += v
		} else {
			bs += v
		}

	}

	if as > bd || ad > bs {
		return true
	}

	return false
}
