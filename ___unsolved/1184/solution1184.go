package solution1184

import "fmt"

// ============================================================================
// 1184. Distance Between Bus Stops
// URL: https://leetcode.com/problems/distance-between-bus-stops/
// ============================================================================

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	ans := 0

	sum1 := 0
	sum2 := 0
	minval1 := 0
	minval2 := 0
	a := start
	b := start
	stop1 := false
	stop2 := false
	for {
		v1 := distance[a]
		v2 := distance[b]
		if !stop1 && a == destination {
			minval1 = sum1
			stop1 = true
		}
		if !stop2 && b == destination {
			minval2 = sum2
			stop2 = true
		}
		if !stop1 && a != destination {
			//fmt.Printf("v1: %d, stop1: %t, minval1: %d\n", v1, stop1, minval1)
			sum1 += v1
		}
		if !stop2 && b != destination {
			sum2 += v2
		}
		fmt.Printf("v2: %d, stop2: %t, minval2: %d\n", v2, stop2, minval2)
		if stop1 && stop2 {
			break
		}

		a++
		if a > len(distance) {
			a = 0
		}
		b--
		if b < 0 {
			b = len(distance) - 1
		}
	}

	ans = min(minval1, minval2)
	fmt.Printf("minval1: %d\n", minval1)
	fmt.Printf("minval2: %d\n", minval2)
	fmt.Printf("sum1: %d\n", sum1)
	fmt.Printf("sum2: %d\n", sum2)

	return ans
}
