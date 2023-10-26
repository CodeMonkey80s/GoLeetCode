package solution2540

func getCommon(nums1 []int, nums2 []int) int {
	freq := make(map[int]int)
	for _, v := range nums1 {
		freq[v]++
	}
	for _, v := range nums2 {
		_, ok := freq[v]
		if ok {
			freq[v]--
		}
	}
	ans := 0
	for {
		for k, v := range freq {
			if v <= ans {
				return k
			}
		}
		ans++
		if ans > 100 {
			break
		}
	}
	return ans
}
