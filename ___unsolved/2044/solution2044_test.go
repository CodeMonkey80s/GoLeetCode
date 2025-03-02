package solution2044

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{3, 1},
		Output: 2,
	},
	{
		Input:  []int{2, 2, 2},
		Output: 7,
	},
	{
		Input:  []int{3, 2, 1, 5},
		Output: 6,
	},
}

func Test_countMaxOrSubsets(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countMaxOrSubsets(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
