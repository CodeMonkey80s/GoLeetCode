package solution836

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	Output bool
}{
	{
		InputA: []int{0, 0, 2, 2},
		InputB: []int{1, 1, 3, 3},
		Output: true,
	},
	{
		InputA: []int{0, 0, 1, 1},
		InputB: []int{1, 0, 2, 1},
		Output: false,
	},
	{
		InputA: []int{0, 0, 1, 1},
		InputB: []int{2, 2, 3, 3},
		Output: false,
	},
}

func Test_isRectangleOverlap(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isRectangleOverlap(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
