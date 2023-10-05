package solution69

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  4,
		Output: 2,
	},
	{
		Input:  8,
		Output: 2,
	},
	// Additional my custom cases
}

func Test_muSqrt(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := mySqrt(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
