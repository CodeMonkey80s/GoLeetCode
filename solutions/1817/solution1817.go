package solution1817

// ============================================================================
// 1817. Finding the Users Active Minutes
// URL: https://leetcode.com/problems/finding-the-users-active-minutes/
// ============================================================================

func findingUsersActiveMinutesV1(logs [][]int, k int) []int {

	var t int
	var id int
	uam := make(map[int]map[int]int)
	for _, log := range logs {
		id = log[0]
		t = log[1]

		if uam[id] == nil {
			uam[id] = make(map[int]int)
		}
		uam[id][t]++
	}

	output := make([]int, k)
	for i := 1; i <= k; i++ {

		var c int
		for _, m := range uam {
			if i == len(m) {
				c++
			}
		}
		output[i-1] = c
	}

	return output
}

func findingUsersActiveMinutesV2_copilot(logs [][]int, k int) []int {

	uam := make(map[int]map[int]struct{})
	for _, log := range logs {
		id := log[0]
		t := log[1]
		if uam[id] == nil {
			uam[id] = make(map[int]struct{})
		}
		uam[id][t] = struct{}{}
	}

	output := make([]int, k)
	for _, times := range uam {
		activeMinutes := len(times)
		if activeMinutes <= k {
			output[activeMinutes-1]++
		}
	}

	return output
}
