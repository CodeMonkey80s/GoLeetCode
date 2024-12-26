package solution2120

// ============================================================================
//
// URL: https://leetcode.com/problems/rearrange-array-elements-by-sign/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2120
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_executeInstructions
	Benchmark_executeInstructions-24    	36075501	        41.21 ns/op	      32 B/op	       1 allocs/op
	PASS
*/

func executeInstructions(n int, startPos []int, s string) []int {

	type robot struct {
		x int
		y int
	}

	rob := robot{
		x: startPos[1],
		y: startPos[0],
	}

	output := make([]int, 0, len(s))

	count := 0

loop:
	for i := 0; i < len(s); i++ {

		if i > 0 {
			output = append(output, count)
		}

		count = 0
		instructions := s[i:]
		rob.x = startPos[1]
		rob.y = startPos[0]

		for _, dir := range instructions {
			switch {
			case dir == 'U':
				if rob.y == 0 {
					continue loop
				}
				rob.y--
				count++
			case dir == 'D':
				if rob.y == n-1 {
					continue loop
				}
				rob.y++
				count++
			case dir == 'R':
				if rob.x == n-1 {
					continue loop
				}
				rob.x++
				count++
			case dir == 'L':
				if rob.x == 0 {
					continue loop
				}
				rob.x--
				count++
			}
		}
	}

	output = append(output, count)

	return output

}
