package solution1710

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA [][]int
	InputB int
	Output int
}{
	{
		InputA: [][]int{{1, 3}, {2, 2}, {3, 1}},
		InputB: 4,
		Output: 8,
	},
	{
		InputA: [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
		InputB: 10,
		Output: 91,
	},
}

func Test_maximumUnits(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maximumUnits(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
