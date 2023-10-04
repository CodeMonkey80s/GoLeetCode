package solution796

// ============================================================================
// 796. Rotate String
// URL: https://leetcode.com/problems/rotate-string/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_rotateString_append-24    	52155769	        26.53 ns/op	      16 B/op	       2 allocs/op
	Benchmark_rotateString_copy-24      	40071675	        29.67 ns/op	      16 B/op	       2 allocs/op
	PASS

*/

func rotateString_append(s string, goal string) bool {
	sl := make([]byte, 0, len(s)+1)
	sl = append(sl, s...)
	i := 0
	for {
		sl = append(sl[1:], sl[0:1]...)
		if string(sl) == goal {
			return true
		}
		i++
		if i > len(s)-1 {
			break
		}
	}
	return false
}

func rotateString_copy(s string, goal string) bool {
	sl := make([]byte, len(s))
	copy(sl, []byte(s))
	i := 0
	ch := byte(0)
	for {
		ch = sl[0]
		sl = sl[1:]
		sl = append(sl, ch)
		if string(sl) == goal {
			return true
		}
		i++
		if i > len(s)-1 {
			break
		}
	}
	return false
}
