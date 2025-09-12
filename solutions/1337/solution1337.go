package solution1337

import "slices"

// ============================================================================
// 1337. The K Weakest Rows in Matrix
// URL: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1337
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_KWeakestRows
	Benchmark_KWeakestRows-24    	19078605	        54.71 ns/op	     104 B/op	       2 allocs/op
	PASS
*/

func kWeakestRows(mat [][]int, k int) []int {
	type row struct {
		index    int
		soldiers int
	}

	rows := make([]row, 0, len(mat))
	for y := 0; y < len(mat); y++ {
		r := row{
			index:    y,
			soldiers: 0,
		}
		for x := 0; x < len(mat[y]); x++ {
			if mat[y][x] == 1 {
				r.soldiers++
			} else {
				break
			}
		}
		rows = append(rows, r)
	}

	slices.SortFunc(rows, func(a, b row) int {
		switch {
		case a.soldiers < b.soldiers:
			return -1
		case a.soldiers == b.soldiers && a.index < b.index:
			return -1
		default:
			return 0
		}
	})

	output := make([]int, 0, k)
	for i := 0; i < k; i++ {
		output = append(output, rows[i].index)
	}

	return output
}
