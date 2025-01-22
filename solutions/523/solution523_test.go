package solution523

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output bool
}{
	{
		InputA: []int{23, 2, 4, 6, 7},
		InputB: 6,
		Output: true,
	},
	{
		InputA: []int{23, 2, 6, 4, 7},
		InputB: 6,
		Output: true,
	},
	{
		InputA: []int{23, 2, 6, 4, 7},
		InputB: 13,
		Output: false,
	},
	{
		InputA: []int{23, 2, 4, 6, 6},
		InputB: 7,
		Output: true,
	},
}

func Test_checkSubarraySum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := checkSubarraySum(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
