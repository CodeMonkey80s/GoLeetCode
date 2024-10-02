package soluton3300

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{10, 12, 13, 14},
		Output: 1,
	},
	{
		Input:  []int{1, 2, 3, 4},
		Output: 1,
	},
	{
		Input:  []int{999, 19, 199},
		Output: 10,
	},
}

func Test_minElement(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minElement(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
