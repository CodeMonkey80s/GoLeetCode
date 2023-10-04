package solution190

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  uint32
	Output uint32
}{
	// Mandatory Test Cases
	{
		Input:  43261596,
		Output: 964176192,
	},
	{
		Input:  4294967293,
		Output: 3221225471,
	},
	// Additional my custom cases
}

func Test_reverseBits(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reverseBits(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
