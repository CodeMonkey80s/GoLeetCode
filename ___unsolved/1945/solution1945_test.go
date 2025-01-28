package solution1945

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB int
	Output int
}{
	{
		InputA: "iiii",
		InputB: 1,
		Output: 36,
	},
	{
		InputA: "leetcode",
		InputB: 2,
		Output: 6,
	},
	{
		InputA: "zbax",
		InputB: 2,
		Output: 8,
	},
}

func Test_getLucky(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getLucky(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
