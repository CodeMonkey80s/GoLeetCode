package solution1652

func decrypt(code []int, k int) []int {

	output := make([]int, len(code))

	var sum int
	for i := 0; i < len(code); i++ {

		sum = 0
		switch {
		case k > 0:
			for j := 1; j <= k; j++ {
				idx := (i + j) % len(code)
				sum += code[idx]
			}
		case k < 0:
			for j := -1; j >= k; j-- {
				idx := (i + j) % len(code)
				if idx < 0 {
					idx = len(code) + idx
				}
				sum += code[idx]
			}
		}

		output[i] = sum
	}

	return output
}
