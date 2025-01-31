package solution1394

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{2, 2, 3, 4},
		Output: 2,
	},
	{
		Input:  []int{1, 2, 2, 3, 3, 3},
		Output: 3,
	},
	{
		Input:  []int{2, 2, 2, 3, 3},
		Output: -1,
	},
}

func Test_findLucky(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findLucky(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
