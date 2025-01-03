package solution20

// ============================================================================
// 20. Valid Parentheses
// URL: https://leetcode.com/problems/valid-parentheses/
// ============================================================================

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/20---Valid-Parentheses
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_isValid
	Benchmark_isValid-24    	100000000	        11.88 ns/op	       2 B/op	       1 allocs/op
	PASS
*/

func isValid(s string) bool {
	length := len(s)
	if length <= 1 || length > 10_000 || length%2 == 1 {
		return false
	}
	var a, b byte
	var l int
	st := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		a = s[i]
		if a == '(' || a == '[' || a == '{' {
			st = append(st, a)
			continue
		}
		l = len(st)
		if i == 0 || l == 0 {
			return false
		}
		if l > 0 {
			b = st[l-1]
			switch {
			case a == ')' && b == '(':
				st = append(st[:l-1], st[l:]...)
			case a == ']' && b == '[':
				st = append(st[:l-1], st[l:]...)
			case a == '}' && b == '{':
				st = append(st[:l-1], st[l:]...)
			default:
				return false
			}
		}
	}
	return len(st) == 0
}

/*
func isValid(s string) bool {
	length := len(s)
	if length <= 1 || length > 10_000 || length%2 == 1 {
		return false
	}
	var a, b byte
	m := map[byte]int{
		'(': 0,
		'[': 0,
		'{': 0,
		')': 0,
		']': 0,
		'}': 0,
	}
	for i := 0; i < length; i++ {
		a = s[i]
		fmt.Printf("i: %d, %c == %c\n", i, a, b)
		if a == '(' || a == '[' || a == '{' {
			m[a]++
			b = a
			fmt.Println("add")
			continue
		}
		if a == ')' && b == '(' ||
			a == ']' && b == '[' ||
			a == '}' && b == '{' {
			m[b]--
			fmt.Println("del")
		}
		b = a
	}
	fmt.Printf("Map (output): %v\n", m)
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
*/

/*
func isValid(s string) bool {
	length := len(s)
	if length < 1 || length > 10_000 {
		return false
	}
	var a, b byte
	var d int
	for i := 0; i < length/2; i++ {
		a = s[d]
		b = s[d+1]
		fmt.Printf("%d, %c = %c\n", d, a, b)
		if a == '(' && b != ')' ||
			a == '[' && b != ']' ||
			a == '{' && b != '}' {
			return false
		}
		d += 2
	}
	return true
}
*/

/*
func isValid_One(s string) bool {
	length := len(s)
	if length < 1 {
		return false
	}
	m := map[byte]int{
		'(': 0,
		'[': 0,
		'{': 0,
	}
	for i := 0; i < length; i++ {
		ch := s[i]
		switch ch {
		case '(':
			m[ch]++
		case '[':
			m[ch]++
		case '{':
			m[ch]++
		case ')':
			m['(']--
		case ']':
			m['[']--
		case '}':
			m['{']--
		}
	}
	fmt.Printf("map: %v\n", m)
	if m['('] == 0 && m['['] == 0 && m['{'] == 0 {
		return true
	}
	return false
}
*/
