package solution3159

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {

	/*
			InputA: []int{1, 3, 1, 7},
			InputB: []int{1, 3, 2, 4},
			InputC: 1,
			Output: []int{0, -1, 2, -1},

			1: 0
		    3: 1
		    7: 3
	*/

	var c int
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n > x {
			nums[n] = i
			continue
		}
		if c == 0 {
			nums[n] = i
			c++
		}
	}

	fmt.Printf("nums: %+v\n", nums)

	for i, n := range queries {
		if x == n {
			queries[i] = nums[n]
		} else {
			queries[i] = -1
		}
	}

	return queries
}

func occurrencesOfElementV2(nums []int, queries []int, x int) []int {

	freq := make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		_, ok := freq[n]
		if !ok {
			freq[n] = make([]int, 0)
		}
		freq[n] = append(freq[n], i)
	}

	var output []int
	for i := 0; i < len(queries); i++ {
		n := queries[i] - 1
		indexes, ok := freq[x]
		if !ok {
			output = append(output, -1)
			continue
		}

		if n < len(indexes) {
			output = append(output, indexes[n])
			continue
		}

		if n >= len(indexes) {
			output = append(output, -1)
			continue
		}
	}

	return output
}
