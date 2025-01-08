package solution1413

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{-3, 2, -3, 4, 2},
		Output: 5,
	},
	{
		Input:  []int{1, 2},
		Output: 1,
	},
	{
		Input:  []int{1, -2, -3},
		Output: 5,
	},
	{
		Input:  []int{-12},
		Output: 13,
	},
}

func Test_minStartValue(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minStartValue(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
