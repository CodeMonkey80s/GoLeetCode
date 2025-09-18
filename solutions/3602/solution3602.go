package solution3602

import "slices"

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3602
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_concatHex36
	Benchmark_concatHex36-24    	41297694	        25.99 ns/op	       5 B/op	       1 allocs/op
	PASS
*/

func concatHex36(n int) string {
	n2 := n * n
	n3 := n2 * n

	hexadec := []byte("0123456789ABCDEF")
	hexatri := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	var s1 []byte
	for n2 > 0 {
		d := n2 % 16
		s1 = append(s1, hexadec[d])
		n2 /= 16
	}

	var s2 []byte
	for n3 > 0 {
		d := n3 % 36
		s2 = append(s2, hexatri[d])
		n3 /= 36
	}

	slices.Reverse(s1)
	slices.Reverse(s2)

	return string(s1) + string(s2)
}
