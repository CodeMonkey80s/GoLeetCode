package solution682

import (
	"strconv"
)

// ============================================================================
// 682. Baseball Game
// URL: https://leetcode.com/problems/baseball-game/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/682---Baseball-Game
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_calPoints
	Benchmark_calPoints-24    	30467395	        33.43 ns/op	      48 B/op	       1 allocs/op
	PASS
*/

func calPoints(operations []string) int {

	scores := make([]int, 0, len(operations))

	for _, op := range operations {
		switch {
		case op == "C":
			scores = scores[:len(scores)-1]
		case op == "D":
			scores = append(scores, scores[len(scores)-1]*2)
		case op == "+":
			scores = append(scores, scores[len(scores)-2]+scores[len(scores)-1])
		default:
			v, _ := strconv.Atoi(op)
			scores = append(scores, v)
		}
	}

	score := 0
	for _, s := range scores {
		score += s
	}

	return score
}
