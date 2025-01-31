package solution1394

func findLucky(arr []int) int {

	freq := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		freq[arr[i]]++
	}

	lucky := -1
	for number, count := range freq {
		if number == count {
			lucky = max(lucky, number)
		}
	}

	return lucky
}
