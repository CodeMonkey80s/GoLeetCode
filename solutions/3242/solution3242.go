package solution3242

// ============================================================================
// 3242. Design Neighbor Sum Service
// URL: https://leetcode.com/problems/design-neighbor-sum-service/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3242---Design-Neighbor-Sum-Service
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	BenchmarkGetFinalState
	BenchmarkGetFinalState-24    	86156235	        13.78 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

type NeighborSum struct {
	grid [][]int
	n    int
}

func Constructor(grid [][]int) NeighborSum {
	return NeighborSum{
		grid: grid,
		n:    len(grid),
	}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	sum := 0
loop:
	for y := 0; y < this.n; y++ {
		for x := 0; x < this.n; x++ {
			if this.grid[y][x] == value {
				if y-1 >= 0 {
					sum += this.grid[y-1][x]
				}
				if x-1 >= 0 {
					sum += this.grid[y][x-1]
				}
				if y+1 < this.n {
					sum += this.grid[y+1][x]
				}
				if x+1 < this.n {
					sum += this.grid[y][x+1]
				}
				break loop
			}
		}
	}

	return sum
}

func (this *NeighborSum) DiagonalSum(value int) int {
	sum := 0
loop:
	for y := 0; y < this.n; y++ {
		for x := 0; x < this.n; x++ {
			if this.grid[y][x] == value {
				if x-1 >= 0 && y-1 >= 0 {
					sum += this.grid[y-1][x-1]
				}
				if x+1 < this.n && y-1 >= 0 {
					sum += this.grid[y-1][x+1]
				}
				if x-1 >= 0 && y+1 < this.n {
					sum += this.grid[y+1][x-1]
				}
				if x+1 < this.n && y+1 < this.n {
					sum += this.grid[y+1][x+1]
				}
				break loop
			}
		}
	}

	return sum
}
