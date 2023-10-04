package solution27

// ============================================================================
// 27. Remove Element
// URL: https://leetcode.com/problems/remove-element/
// ============================================================================

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	idx := -1
	i := 0
	c := 0
	for {
		if nums[i] == val {
			nums[i] = -1
			idx = i
			c++
		}
		if idx != -1 {
			for j := idx; j < n; j++ {
				if j < n-1 {
					nums[j] = nums[j+1]
				}
			}
			nums[n-1] = -1
			idx = -1
			continue
		}
		i++
		if i >= n {
			break
		}
	}
	return n - c
}

func removeElement_stdlib(nums []int, val int) int {
loop:
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] == val {
			if i == 0 {
				nums = nums[1:]
				goto loop
			} else if i < n-1 {
				nums = append(nums[:i], nums[i+1:]...)
				goto loop
			} else {
				nums = nums[:i]
				goto loop
			}
		}
	}
	return len(nums)
}
