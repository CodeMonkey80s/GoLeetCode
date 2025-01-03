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
	Benchmark_maxNumberOfBalloons-24                  	84417007	        14.55 ns/op	       0 B/op	       0 allocs/op
	Benchmark_maxNumberOfBalloons_first_attempt-24    	 4507036	       266.8 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func maxNumberOfBalloons(text string) int {
	if len(text) < 7 {
		return 0
	}
	m := [5]int{0, 0, 0, 0, 0}
	for _, ch := range text {
		switch ch {
		case 'b':
			m[0]++
		case 'a':
			m[1]++
		case 'l':
			m[2]++
		case 'o':
			m[3]++
		case 'n':
			m[4]++
		}
	}
	num := 0
outer:
	for {
		m[0] -= 1
		m[1] -= 1
		m[2] -= 2
		m[3] -= 2
		m[4] -= 1
		switch {
		case m[0] < 0:
			break outer
		case m[1] < 0:
			break outer
		case m[2] < 0:
			break outer
		case m[3] < 0:
			break outer
		case m[4] < 0:
			break outer
		}
		num++
	}
	return num
}

func maxNumberOfBalloons_first_attempt(text string) int {
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
