package solution724

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 7, 3, 6, 5, 6},
		Output: 3,
	},
	{
		Input:  []int{1, 2, 3},
		Output: -1,
	},
	{
		Input:  []int{2, 1, -1},
		Output: 0,
	},
	{
		Input:  []int{-1, -1, -1, -1, -1, -1},
		Output: -1,
	},
	{
		Input:  []int{0},
		Output: 0,
	},
}

func Test_pivotIndex(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := pivotIndexV2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
