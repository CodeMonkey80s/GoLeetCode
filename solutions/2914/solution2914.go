package solution2914

func minChanges(s string) int {

	var n int
	for i := 0; i < len(s)-1; i += 2 {
		switch {
		case s[i] == '1' && s[i+1] == '0':
			n++
		case s[i] == '0' && s[i+1] == '1':
			n++
		}
	}

	return n
}
