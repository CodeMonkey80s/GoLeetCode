package solution3168

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "EEEEEEE",
		Output: 7,
	},
	{
		Input:  "ELELEEL",
		Output: 2,
	},
}

func Test_clearDigits(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %q, Output: %d\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minimumChairs(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %d but we got %d", tc.Output, output)
			}
		})
	}
}
