package solution3561

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3461
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_makeGood
	Benchmark_makeGood-24    	153518443	         7.817 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func hasSameDigits(s string) bool {
	st := []byte(s)
loop:
	for i := 0; i < len(st); i++ {
		a := i
		b := i + 1
		n := (st[a] - '0' + st[b] - '0') % 10
		st[i] = n + '0'
		a++
		b++
		if b > len(st)-1 {
			st = st[:len(st)-1]
			if len(st) == 2 {
				break
			}
			goto loop
		}
	}
	return st[0] == st[1]
}
