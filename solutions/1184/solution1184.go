package solution1184

// ============================================================================
// 1184. Distance Between Bus Stops
// URL: https://leetcode.com/problems/distance-between-bus-stops/
// ============================================================================

func distanceBetweenBusStopsV2(distance []int, start int, destination int) int {

	var dist1 int
	for i := start; i != destination; i = (i + 1) % len(distance) {
		dist1 += distance[i]
	}

	var dist2 int
	for i := destination; i != start; i = (i + 1) % len(distance) {
		dist2 += distance[i]
	}

	return min(dist1, dist2)
}
