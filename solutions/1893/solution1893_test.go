package solution1893

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA [][]int
	InputB int
	InputC int
	Output bool
}{
	{
		InputA: [][]int{{1, 2}, {3, 4}, {5, 6}},
		InputB: 2,
		InputC: 5,
		Output: true,
	},
	{
		InputA: [][]int{{1, 10}, {10, 20}},
		InputB: 21,
		InputC: 21,
		Output: false,
	},
}

func Test_isCovered(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isCovers(tc.InputA, tc.InputB, tc.InputC)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
