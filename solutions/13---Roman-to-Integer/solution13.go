package solution13

// ============================================================================
// 13. Roman to Integer
// URL: https://leetcode.com/problems/roman-to-integer/
// ============================================================================

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	length := len(s)
	if length < 1 {
		return 0
	}
	var a, b byte
	var sum int
	for i := 0; i < length; i++ {
		a = s[i]
		if i < length-1 {
			b = s[i+1]
		}
		switch {
		case a == 'I' && b == 'V':
			sum += 4
			i++
		case a == 'I' && b == 'X':
			sum += 9
			i++
		case a == 'X' && b == 'L':
			sum += 40
			i++
		case a == 'X' && b == 'C':
			sum += 90
			i++
		case a == 'C' && b == 'D':
			sum += 400
			i++
		case a == 'C' && b == 'M':
			sum += 900
			i++
		default:
			sum += m[a]
		}
	}
	return sum
}
