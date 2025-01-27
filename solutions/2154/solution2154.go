package solution2154

func findFinalValue(nums []int, original int) int {

	var finalValue int
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	for {
		_, ok := freq[original]
		if !ok {
			finalValue = original
			break
		}

		original *= 2
	}

	return finalValue
}
