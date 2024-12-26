package solution861

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output int
}{
	{
		Input:  [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}},
		Output: 39,
	},
	{
		Input:  [][]int{{0, 1, 1}, {1, 1, 1}, {0, 1, 0}},
		Output: 18,
	},
	{
		Input:  [][]int{{0}},
		Output: 1,
	},
}

func Test_matrixScore(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := matrixScore(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
