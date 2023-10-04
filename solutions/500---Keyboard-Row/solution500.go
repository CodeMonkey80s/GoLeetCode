package solution500

// ============================================================================
// 500. Keyboard Row
// URL: https://leetcode.com/problems/keyboard-row/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findWords-24    	  922678	      1354 ns/op	    1609 B/op	       4 allocs/op
	PASS

*/

func findWords(words []string) []string {
	m := map[byte]int{
		'q': 1,
		'w': 1,
		'e': 1,
		'r': 1,
		'y': 1,
		'u': 1,
		'i': 1,
		'o': 1,
		'p': 1,
		'a': 2,
		's': 2,
		'd': 2,
		'f': 2,
		'g': 2,
		'h': 2,
		'j': 2,
		'k': 2,
		'l': 2,
		'z': 3,
		'x': 3,
		'c': 3,
		'v': 3,
		'b': 3,
		'n': 3,
		'm': 3,
		'Q': 1,
		'W': 1,
		'E': 1,
		'R': 1,
		'Y': 1,
		'U': 1,
		'I': 1,
		'O': 1,
		'P': 1,
		'A': 2,
		'S': 2,
		'D': 2,
		'F': 2,
		'G': 2,
		'H': 2,
		'J': 2,
		'K': 2,
		'L': 2,
		'Z': 3,
		'X': 3,
		'C': 3,
		'V': 3,
		'B': 3,
		'N': 3,
		'M': 3,
	}
	row := 0
	output := []string{}
outer:
	for _, word := range words {
		row = 0
		for _, rune := range word {
			ch := byte(rune)
			v, ok := m[ch]
			switch {
			case ok && row == 0:
				row = v
			case ok && row != 0 && row != v:
				continue outer
			}
		}
		if row != 0 {
			output = append(output, word)
		}
	}
	return output
}
