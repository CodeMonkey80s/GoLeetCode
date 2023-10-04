package solution771

import (
	"bytes"
	"strings"
)

// ============================================================================
// 771. Jewels and Stones
// URL: https://leetcode.com/problems/jewels-and-stones/
// ============================================================================

/*
	$ go test -bench=. -benchmem -benchtime=5s
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_numJewelsInStones_bytes-24      	557335704	        10.78 ns/op	       0 B/op	       0 allocs/op
	Benchmark_numJewelsInStones_strings-24    	786849105	         7.617 ns/op	   0 B/op	       0 allocs/op
	Benchmark_numJewelsInStones_my-24         	1000000000	         4.006 ns/op	   0 B/op	       0 allocs/op
	PASS
*/

func numJewelsInStones_bytes(jewels string, stones string) int {
	n := 0
	j := []byte(jewels)
	s := []byte(stones)
	for _, v := range j {
		n += bytes.Count(s, []byte{v})
	}
	return n
}

func numJewelsInStones_strings(jewels string, stones string) int {
	n := 0
	for _, v := range jewels {
		n += strings.Count(stones, string(v))
	}
	return n
}

func numJewelsInStones(jewels string, stones string) int {
	n := 0
	for _, c1 := range jewels {
		for _, c2 := range stones {
			if c1 == c2 {
				n++
			}
		}
	}
	return n
}
