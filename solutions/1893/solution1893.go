package solution1893

func isCovers(ranges [][]int, left int, right int) bool {

	var c int

loop:
	for i := left; i <= right; i++ {
		for _, nums := range ranges {
			if nums[0] <= i && i <= nums[1] {
				c++
				continue loop
			}
		}
	}

	return c == right-left+1
}
