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
	sl := make([]byte, len(s))
	copy(sl, s)
	var a, b byte
	i := 0
	m := len(sl)
	for {
		if i+1 > m-1 {
			break
		}
		a = sl[i]
		b = sl[i+1]
		if a == b {
			sl = append(sl[:i], sl[i+2:]...)
			i = 0
			m = len(sl)
			continue
		}
		i++
	}
	return string(sl)
}

func removeDuplicates_string(s string) string {
	stack := make([]rune, 0)
	for _, elem := range s {
		if len(stack) != 0 && stack[len(stack)-1] == elem {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, elem)
		}
	}
	return string(stack)
}
