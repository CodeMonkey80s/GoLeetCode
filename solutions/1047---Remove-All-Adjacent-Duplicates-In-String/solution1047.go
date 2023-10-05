package solution1047

// ============================================================================
// 1047. Remove All Adjacent Duplicates In String
// URL: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
// ============================================================================

/*
	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_removeDuplicates-24           	85594324	        16.28 ns/op	       8 B/op	       1 allocs/op
	Benchmark_removeDuplicates_string-24    	48264585	        27.28 ns/op	       8 B/op	       1 allocs/op
	PASS
*/

func removeDuplicates(s string) string {
	stack := make([]byte, len(s))
	copy(stack, s)
	var a, b byte
	i := 0
	m := len(stack)
	for {
		if i+1 > m-1 {
			break
		}
		a = stack[i]
		b = stack[i+1]
		if a == b {
			stack = append(stack[:i], stack[i+2:]...)
			i = 0
			m = len(stack)
			continue
		}
		i++
	}
	return string(stack)
}

func removeDuplicatesStack(s string) string {
	stack := make([]rune, 0)
	for _, char := range s {
		if len(stack) != 0 && stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return string(stack)
}
