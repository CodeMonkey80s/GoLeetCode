package solution521

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output int
}{
	{
		InputA: "aba",
		InputB: "cdc",
		Output: 3,
	},
	{
		InputA: "aaa",
		InputB: "bbb",
		Output: 3,
	},
	{
		InputA: "aaa",
		InputB: "aaa",
		Output: -1,
	},
}

func Test_findLUSlength(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findLUSlength(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
