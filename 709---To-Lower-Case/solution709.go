package solution709

// ============================================================================
// 709. To Lower Case
// URL: https://leetcode.com/problems/to-lower-case
// ============================================================================

/*
	$ go test -bench=. -benchmem -benchtime=5s
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_toLowerCase_my-24     	434713898	        14.76 ns/op	       5 B/op	       1 allocs/op
	Benchmark_toLowerCase_std-24    	259508497	        22.53 ns/op	       5 B/op	       1 allocs/op
	PASS
*/

func toLowerCase(s string) string {
	sl := make([]byte, 0, len(s))
	for _, v := range s {
		switch {
		case v >= 'A' && v <= 'Z':
			v += 32
		}
		sl = append(sl, byte(v))
	}
	return string(sl)
}
