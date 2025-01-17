package solution2908

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{8, 6, 1, 5, 3},
		Output: 9,
	},
	{
		Input:  []int{5, 4, 8, 7, 10, 2},
		Output: 13,
	},
	{
		Input:  []int{6, 5, 4, 3, 4, 5},
		Output: -1,
	},
}

func Test_minimumSum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minimumSum(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
