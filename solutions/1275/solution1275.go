package solution1275

// ============================================================================
// 1275. Find Winner on a Tic Tac Toe Game
// URL: https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
// ============================================================================

/*

	go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/1275---Find-Winner-on-a-Tic-Tac-Toe-Game
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_tictactoe-24    	19551152	        61.16 ns/op	       0 B/op	       0 allocs/op
	PASS

*/

func tictactoe(moves [][]int) string {
	boardA := uint(0)
	boardB := uint(0)
	victories := [8]uint{
		0b_111000000,
		0b_000111000,
		0b_000000111,
		0b_100100100,
		0b_010010010,
		0b_001001001,
		0b_100010001,
		0b_001010100,
	}
	a, b, x, y := 0, 0, 0, 0
	idx := 0
	for i, pair := range moves {
		x = pair[0]
		y = pair[1]
		idx = y*3 + x
		if i%2 == 0 {
			boardA |= 1 << idx
		} else {
			boardB |= 1 << idx
		}
	}
	for _, mask := range victories {
		a = 0
		b = 0
		for i := 0; i < 9; i++ {
			v := mask & (1 << i)
			va := boardA & (1 << i)
			vb := boardB & (1 << i)
			if v > 0 && va > 0 && v == va {
				a++
			}
			if v > 0 && vb > 0 && v == vb {
				b++
			}
		}
		if a == 3 {
			return "A"
		}
		if b == 3 {
			return "B"
		}
	}
	if len(moves) < 9 {
		return "Pending"
	}
	return "Draw"
}
