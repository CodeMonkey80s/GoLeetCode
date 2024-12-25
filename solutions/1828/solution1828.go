package solution1828

// ============================================================================
// 1828. Queries on Number of Points Inside a Circle
// URL: https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/
// ============================================================================

func countPoints(points [][]int, queries [][]int) []int {
	output := make([]int, 0, len(queries))
	for _, q := range queries {
		count := 0
		for _, p := range points {
			if p[0] > q[0]+q[2] && p[1] > q[1]+q[2] {
				continue
			}
			d := (q[0]-p[0])*(q[0]-p[0]) + (q[1]-p[1])*(q[1]-p[1])
			r := q[2] * q[2]
			if d <= r {
				count++
			}
		}
		output = append(output, count)
	}

	return output
}
