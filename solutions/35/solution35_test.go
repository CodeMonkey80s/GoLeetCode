package solution35

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Target int
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  []int{1, 3, 5, 6},
		Target: 5,
		Output: 2,
	},
	{
		Input:  []int{1, 3, 5, 6},
		Target: 2,
		Output: 1,
	},
	{
		Input:  []int{1, 3, 5, 6},
		Target: 7,
		Output: 4,
	},
	// Additional my custom cases
	{
		Input:  []int{1, 3, 5, 6},
		Target: 0,
		Output: 0,
	},
	{
		Input:  []int{1, 3},
		Target: 2,
		Output: 1,
	},
	{
		Input:  []int{1, 3, 5},
		Target: 4,
		Output: 2,
	},
	{
		Input:  []int{1, 3, 5},
		Target: 5,
		Output: 2,
	},
	{
		Input:  []int{1, 2, 3, 4, 5, 10},
		Target: 2,
		Output: 1,
	},
	{
		Input:  []int{1, 2, 4, 6, 8, 9, 10},
		Target: 10,
		Output: 6,
	},
}

func Test_searchInsert(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Target: %v, Output: %v\n", tc.Input, tc.Target, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := searchInsert(tc.Input, tc.Target)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_searchInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = searchInsert(testCases[0].Input, testCases[0].Target)
	}
}
