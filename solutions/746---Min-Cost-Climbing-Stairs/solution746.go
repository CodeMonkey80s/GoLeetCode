package solution746

// ============================================================================
// 746. Min Cost Climbing Stairs
// URL: https://leetcode.com/problems/min-cost-climbing-stairs/
// ============================================================================

//
// The IF/ELSE calculating min can be replaced with min() in Go v1.21+
//

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	size := len(cost)
	for i := size - 3; i > -1; i-- {
		if cost[i+1] > cost[i+2] {
			cost[i] += cost[i+2]
		} else {
			cost[i] += cost[i+1]
		}
	}
	if cost[0] > cost[1] {
		return cost[1]
	}
	return cost[0]
}
