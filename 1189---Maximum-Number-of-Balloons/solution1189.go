package solution1189

// ============================================================================
// 1189. Maximum Number of Balloons
// URL: https://leetcode.com/problems/maximum-number-of-balloons/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_maxNumberOfBalloons-24    	 4035127	       294.5 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func maxNumberOfBalloons(text string) int {
	if len(text) < 7 {
		return 0
	}
	m := map[byte]int{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}
	for _, ru := range text {
		ch := byte(ru)
		if ch == 'b' || ch == 'a' || ch == 'l' || ch == 'o' || ch == 'n' {
			m[ch]++
		}
	}
	num := 0
outer:
	for {
		m['b'] -= 1
		m['a'] -= 1
		m['l'] -= 2
		m['o'] -= 2
		m['n'] -= 1
		switch {
		case m['b'] < 0:
			break outer
		case m['a'] < 0:
			break outer
		case m['l'] < 0:
			break outer
		case m['o'] < 0:
			break outer
		case m['n'] < 0:
			break outer
		}
		num++
	}
	return num
}
