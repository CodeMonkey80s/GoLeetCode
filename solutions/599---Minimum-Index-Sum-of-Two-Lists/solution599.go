package solution599

// ============================================================================
// 599. Minimum Index Sum of Two Lists
// URL: https://leetcode.com/problems/minimum-index-sum-of-two-lists/
// ============================================================================

/*

 ***** Defunc, maybe I'll revisit later...

func findRestaurant(list1 []string, list2 []string) []string {
	output := []string{}
	var maxval int = -1
	for i, s1 := range list1 {
		for j, s2 := range list2 {
			if s1 == s2 {
				val := i + j
				if maxval == -1 && maxval < val {
					maxval = val
				}
				if maxval > val {
					break
				}
				if maxval != -1 && val == maxval {
					output = append(output, s1)
				}
			}
		}
	}
	return output
}

*/

func findRestaurant_first_attempt(list1 []string, list2 []string) []string {
	count := len(list1)
	if len(list2) > count {
		count = len(list2)
	}
	m := make(map[string]int, count)
	output := []string{}
	var minval int = 1<<63 - 1
	for i, s1 := range list1 {
		_, ok1 := m[s1]
		if !ok1 {
			m[s1] = -1
		}
		for j, s2 := range list2 {
			switch {
			case s1 == s2:
				m[s1] = 0
				m[s1] += i
				m[s2] += j
				if m[s2] < minval {
					minval = m[s2]
				}
			}
		}
	}
	for k, v := range m {
		if v != -1 {
			if v == minval {
				output = append(output, k)
			}
		}
	}
	return output
}
