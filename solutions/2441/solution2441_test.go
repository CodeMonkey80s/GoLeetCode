package solution2441

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{-1, 2, -3, 3},
		Output: 3,
	},
	{
		Input:  []int{-1, 10, 6, 7, -7, 1},
		Output: 7,
	},
	{
		Input:  []int{-10, 8, 6, 7, -2, -3},
		Output: -1,
	},
	{
		Input:  []int{-30, 34, 1, 32, 26, -9, -30, 22, -49, 29, 48, 47, 38, 4, 43, 12, -1, -8, 11, -37, 32, 40, 9, 15, -34, -34, -16, -5, 26, -44, -36, -13, -16, 10, 39, -17, -22, 17, -16},
		Output: 34,
	},
}

func Test_findMax(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findMax(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
