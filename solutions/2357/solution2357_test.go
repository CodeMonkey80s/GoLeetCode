package solution2357

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 5, 0, 3, 5},
		Output: 3,
	},
	{
		Input:  []int{0},
		Output: 0,
	},
	{
		Input:  []int{1},
		Output: 1,
	},
}

func Test_minimumOperations(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minimumOperations(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
