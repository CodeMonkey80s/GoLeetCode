package solution2697

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "egcfe",
		Output: "efcfe",
	},
	{
		Input:  "abcd",
		Output: "abba",
	},
	{
		Input:  "seven",
		Output: "neven",
	},
}

func Test_makeSmallestPalindrome(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := makeSmallestPalindrome(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
