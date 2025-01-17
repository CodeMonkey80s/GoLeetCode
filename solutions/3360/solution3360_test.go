package solution3360

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output bool
}{
	{
		Input:  12,
		Output: true,
	},
	{
		Input:  1,
		Output: false,
	},
	{
		Input:  19,
		Output: false,
	},
}

func Test_canAliceWin(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := canAliceWin(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
