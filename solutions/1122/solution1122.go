package solution1122

import (
	"slices"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {

	freq := make(map[int]int)
	for i := range arr2 {
		freq[arr2[i]] = i
	}

	part1 := make([]int, 0)
	part2 := make([]int, 0)
	for i := 0; i < len(arr1); i++ {
		_, ok := freq[arr1[i]]
		switch {
		case ok:
			part1 = append(part1, arr1[i])
		default:
			part2 = append(part2, arr1[i])
		}
	}

	slices.SortFunc(part1, func(a, b int) int {
		pos1 := freq[a]
		pos2 := freq[b]
		switch {
		case pos1 < pos2:
			return -1
		case pos2 > pos1:
			return 1
		default:
			return 0
		}
	})

	slices.Sort(part2)

	return slices.Concat(part1, part2)
}
