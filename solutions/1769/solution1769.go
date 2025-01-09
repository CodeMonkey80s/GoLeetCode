package solution1769

func minOperations(boxes string) []int {

	prefixSumL := make([]int, len(boxes))
	prefixSumR := make([]int, len(boxes))

	for i := 0; i < len(boxes); i++ {
		for j := i; j < len(boxes); j++ {
			if boxes[j] == '0' {
				continue
			}
			prefixSumR[i] += j - i
		}
	}

	for i := len(boxes) - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if boxes[j] == '0' {
				continue
			}
			prefixSumL[i] += i - j
		}
	}

	var output []int
	for i := range boxes {
		output = append(output, prefixSumL[i]+prefixSumR[i])
	}

	return output
}
