package solution414

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{3, 2, 1},
		Output: 1,
	},
	{
		Input:  []int{1, 2},
		Output: 2,
	},
	{
		Input:  []int{2, 2, 3, 1},
		Output: 1,
	},
	{
		Input:  []int{1, 1, 2},
		Output: 2,
	},
	{
		Input:  []int{1, 1, 1},
		Output: 1,
	},
	{
		Input:  []int{1, 2, 2, 5, 3, 5},
		Output: 2,
	},
}

func Test_thirdMax(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := thirdMax(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
