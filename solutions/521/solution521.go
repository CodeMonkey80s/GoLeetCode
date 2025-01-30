package solution521

func findLUSlength(a string, b string) int {

	if a == b {
		return -1
	}

	return max(len(a), len(b))
}
