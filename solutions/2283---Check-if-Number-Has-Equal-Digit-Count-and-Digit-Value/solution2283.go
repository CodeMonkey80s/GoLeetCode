package solution2283

// ============================================================================
// 2283. Check if Number Has Equal Digit Count and Digit Value
// URL: https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2283---Check-if-Number-Has-Equal-Digit-Count-and-Digit-Value
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_digitCount-24    	267768547	         4.490 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func digitCount(num string) bool {
	freq := make([]byte, 10)
	for _, digit := range num {
		val := byte(digit) - '0'
		freq[val]++
	}
	for i, v := range freq {
		if i < len(num) {
			val := byte(v + '0')
			if i == 0 && val == 0 {
				continue
			}
			if val != num[i] {
				return false
			}
		}
	}
	return true
}
