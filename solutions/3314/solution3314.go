package solution3314

func minBitwiseArray(nums []int) []int {

	find := func(n int) int {
		for i := 0; i < n; i++ {
			v := i | (i + 1)
			if n == v {
				return i
			}
		}

		return -1
	}

	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = find(num)
	}

	return res
}
