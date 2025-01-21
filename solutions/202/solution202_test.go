package solution202

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output bool
}{
	{
		Input:  19,
		Output: true,
	},
	{
		Input:  2,
		Output: false,
	},
}

func Test_isHappy(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isHappyV2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, tc.Input)
			}
		})
	}
}
