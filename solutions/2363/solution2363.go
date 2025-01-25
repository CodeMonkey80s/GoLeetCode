package solution2363

import (
	"slices"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {

	maxItems := max(len(items1), len(items2))
	output := make([][]int, 0, maxItems)

	sums := make(map[int]int)
	for i := 0; i < maxItems; i++ {

		if i < len(items1) {
			value := items1[i][0]
			weight := items1[i][1]
			sums[value] += weight
		}

		if i < len(items2) {
			value := items2[i][0]
			weight := items2[i][1]
			sums[value] += weight
		}
	}

	for value, weight := range sums {
		output = append(output, []int{value, weight})
	}

	slices.SortFunc(output, func(a, b []int) int {
		return a[0] - b[0]
	})

	return output
}
