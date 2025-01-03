package solution2325

// ============================================================================
// 2325. Decode the Message
// URL: https://leetcode.com/problems/decode-the-message/
// ============================================================================

/*
	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2325---Decode-the-Message
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_decodeMessage-24    	 1101666	      1182 ns/op	     290 B/op	       3 allocs/op
	PASS

*/

func decodeMessage(key string, message string) string {
	m := make(map[byte]byte, 26)
	idx := 0
	for _, char := range key {
		if char != ' ' {
			ch := byte(char)
			_, ok := m[ch]
			if !ok {
				m[ch] = byte(idx + 'a')
				idx++
			}
		}
	}
	ans := make([]byte, 0, len(message))
	for _, char := range message {
		if char == ' ' {
			ans = append(ans, ' ')
		} else {
			ch := byte(char)
			ans = append(ans, m[ch])
		}
	}
	return string(ans)
}
