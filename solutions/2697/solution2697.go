package solution2697

func makeSmallestPalindrome(s string) string {

	output := []byte(s)
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			continue
		}
		output[i] = min(s[i], s[j])
		output[j] = output[i]
	}

	return string(output)
}
