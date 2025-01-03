package solution405

import (
	"math"
)

// ============================================================================
// 405. Convert a Number to Hexadecimal
// URL: https://leetcode.com/problems/convert-a-number-to-hexadecimal/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_toHex_my-24    	 3589454	       335.2 ns/op	     360 B/op	       3 allocs/op
	Benchmark_toHex_std-24    	100000000	        12.99 ns/op	       2 B/op	       1 allocs/op

*/

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num += int(math.Pow(2, 32))
	}
	m := map[int]byte{
		0:  '0',
		1:  '1',
		2:  '2',
		3:  '3',
		4:  '4',
		5:  '5',
		6:  '6',
		7:  '7',
		8:  '8',
		9:  '9',
		10: 'a',
		11: 'b',
		12: 'c',
		13: 'd',
		14: 'e',
		15: 'f',
	}
	res := ""
	for num > 0 {
		d := num % 16
		v, ok := m[d]
		if ok {
			res = string(v) + res
		}
		num = int(math.Floor(float64(num) / 16))
	}
	return res
}
