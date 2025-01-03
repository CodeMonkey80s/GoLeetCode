package solution832

import (
	"slices"
)

// ============================================================================
// 832. Flipping an Image
// URL: https://leetcode.com/problems/flipping-an-image/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/832---Flipping-an-Image
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_flipAndInvertImage
	Benchmark_flipAndInvertImage-24    	100000000	        12.18 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func flipAndInvertImage(image [][]int) [][]int {
	for y := 0; y < len(image); y++ {
		slices.Reverse(image[y])
		for x := 0; x < len(image[y]); x++ {
			image[y][x] ^= 1
		}
	}

	return image
}
