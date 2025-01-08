package solution3396

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{

	{
		Input:  []int{1, 2, 3, 4, 2, 3, 3, 5, 7},
		Output: 2,
	},
	{
		Input:  []int{4, 5, 6, 4, 4},
		Output: 2,
	},
	{
		Input:  []int{6, 7, 8, 9},
		Output: 0,
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
