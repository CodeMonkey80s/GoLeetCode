package solution179

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output string
}{
	{
		Input:  []int{10, 2},
		Output: "210",
	},
	{
		Input:  []int{3, 30, 34, 5, 9},
		Output: "9534330",
	},
}

func Test_largestNumber(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := largestNumber(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
