package solution1539

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output int
}{
	{
		InputA: []int{2, 3, 4, 7, 11},
		InputB: 5,
		Output: 9,
	},
	{
		InputA: []int{1, 2, 3, 4},
		InputB: 2,
		Output: 6,
	},
}

func Test_findKthPositive(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findKthPositive(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
