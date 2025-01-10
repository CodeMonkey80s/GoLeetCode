package solution2079

// ============================================================================
// 2079. Watering Plants
// URL: https://leetcode.com/problems/watering-plants/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2079
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_wateringPlantsV1
	Benchmark_wateringPlantsV1-24    	281480796	         4.154 ns/op	       0 B/op	       0 allocs/op
	Benchmark_wateringPlantsV2
	Benchmark_wateringPlantsV2-24    	295546256	         4.055 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func wateringPlantsV2(plants []int, capacity int) int {

	steps := 0
	c := capacity
	for i := 0; i < len(plants); i++ {
		switch {
		case c >= plants[i]:
			c -= plants[i]
			plants[i] = 0
			steps++
		default:
			c = capacity
			c -= plants[i]
			plants[i] = 0
			steps += 1 + (2 * i)
		}
	}

	return steps
}

func wateringPlantsV1(plants []int, capacity int) int {

	d := 1
	x := -1
	c := capacity
	steps := 0

	for {

		x += d
		steps++

		if x+d < 0 {
			c = capacity
			d = -d
			continue
		}

		if c >= plants[x] {
			c -= plants[x]
			plants[x] = 0
		} else {
			d = -d
			continue
		}

		if x+d >= len(plants) {
			break
		}
	}

	return steps
}
