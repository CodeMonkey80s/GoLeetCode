package solution2500

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output int
}{
	{
		Input:  [][]int{{1, 2, 4}, {3, 3, 1}},
		Output: 8,
	},
	{
		Input:  [][]int{{10}},
		Output: 10,
	},
}

func Test_deleteGreatestValue(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := deleteGreatestValue(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
