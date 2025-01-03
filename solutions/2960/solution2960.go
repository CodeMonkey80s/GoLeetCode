package solution2960

// ============================================================================
// 2960. Count Tested Devices After Test Operations
// URL: https://leetcode.com/problems/count-tested-devices-after-test-operations/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2960---Count-Tested-Devices-After-Test-Operations
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_CountTestedDevicesV2
	Benchmark_CountTestedDevicesV2-24    	359160918	         2.953 ns/op	       0 B/op	       0 allocs/op
	Benchmark_CountTestedDevicesV1
	Benchmark_CountTestedDevicesV1-24    	157039726	         7.635 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countTestedDevicesV2(batteryPercentages []int) int {

	n := 0
	for i := range batteryPercentages {
		if batteryPercentages[i] > n {
			n++
		}
	}

	return n
}

func countTestedDevicesV1(batteryPercentages []int) int {

	n := 0
	for i := range batteryPercentages {
		if batteryPercentages[i] == 0 {
			continue
		}

		for j := i + 1; j < len(batteryPercentages); j++ {
			if batteryPercentages[j] > 0 {
				batteryPercentages[j]--
			}
		}

		n++
	}

	return n
}
