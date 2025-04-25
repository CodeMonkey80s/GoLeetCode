package solution3516

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA int
	InputB int
	InputC int
	Output int
}{

	{
		InputA: 2,
		InputB: 7,
		InputC: 4,
		Output: 1,
	},
	{
		InputA: 2,
		InputB: 5,
		InputC: 6,
		Output: 2,
	},
	{
		InputA: 1,
		InputB: 5,
		InputC: 3,
		Output: 0,
	},
}

func Test_findClosest(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findClosest(tc.InputA, tc.InputB, tc.InputC)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
