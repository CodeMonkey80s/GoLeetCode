package solution1544

import "slices"

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1544
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_makeGood
	Benchmark_makeGood-24    	59022217	        19.25 ns/op	       8 B/op	       1 allocs/op
	PASS
*/

func makeGood(s string) string {
	sameChars := func(ch1 byte, ch2 byte) bool {
		if ch1-ch2 == 32 || ch2-ch1 == 32 || ch2+ch1 == 32 {
			return true
		}
		return false
	}

	a := 0
	b := 1
	st := []byte(s)
loop:
	for b <= len(st)-1 {
		if sameChars(st[a], st[b]) {
			st = slices.Delete(st, a, b+1)
			a = 0
			b = 1
			goto loop
		}
		a++
		b++
	}
	return string(st)
}
