package solution3438

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3438
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_makeGood
	Benchmark_makeGood-24    	51707838	        22.57 ns/op	       2 B/op	       1 allocs/op
	PASS
*/

func findValidPair(s string) string {
	var freq [10]int
	for _, ch := range s {
		freq[int(ch-'0')]++
	}

	for i := 0; i < len(s)-1; i++ {
		d1 := s[i]
		d2 := s[i+1]
		switch {
		case d1 == d2:
			continue
		case freq[int(d1-'0')] == int(d1-'0') && freq[int(d2-'0')] == int(d2-'0'):
			return string(d1) + string(d2)
		}
	}

	return ""
}
