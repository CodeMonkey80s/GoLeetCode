package solution2206

func divideArray(nums []int) bool {

	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	for _, count := range freq {
		if count%2 != 0 {
			return false
		}
	}

	return true
}
