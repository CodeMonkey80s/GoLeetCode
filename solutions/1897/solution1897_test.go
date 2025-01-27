package solution1897

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output bool
}{
	{
		Input:  []string{"abc", "aabc", "bc"},
		Output: true,
	},
	{
		Input:  []string{"ab", "a"},
		Output: false,
	},
}

func Test_makeEqual(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := makeEqual(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
