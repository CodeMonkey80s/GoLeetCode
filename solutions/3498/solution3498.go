package solution3498

func reverseDegree(s string) int {
	var degree int
	for i := 0; i < len(s); i++ {
		degree += (123 - int(s[i])) * (i + 1)
	}

	return degree
}
