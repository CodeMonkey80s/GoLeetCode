package solution657

// ============================================================================
// 657. Robot Return to Origin
// URL: https://leetcode.com/problems/robot-return-to-origin/
// ============================================================================

/*

	$ go test -bench=. -benchmem -benchtime=5s
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_judgeCircle_switch-24    	1000000000	         1.141 ns/op	   0 B/op	       0 allocs/op
	Benchmark_judgeCircle_map-24       	79862971	        75.18 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func judgeCircle_switch(moves string) bool {
	x, y := 0, 0
	for _, v := range moves {
		switch v {
		case 'U':
			y--
		case 'D':
			y++
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}

// ***** Proposed by the ChatGPT as "optimization" ... lol...
func judgeCircle_map(moves string) bool {
	x, y := 0, 0
	m := map[byte][2]int{
		'U': {0, -1},
		'D': {0, 1},
		'L': {-1, 0},
		'R': {1, 0},
	}
	mbytes := []byte(moves)
	for _, v := range mbytes {
		move := m[v]
		x += move[0]
		y += move[1]
	}
	return x == 0 && y == 0
}
