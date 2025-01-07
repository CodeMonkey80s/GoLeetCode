package solution3407

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output bool
}{

	{
		InputA: "leetcode",
		InputB: "ee*e",
		Output: true,
	},
	{
		InputA: "car",
		InputB: "c*v",
		Output: false,
	},
	{
		InputA: "luck",
		InputB: "u*",
		Output: true,
	},
	{
		InputA: "xks",
		InputB: "s*",
		Output: true,
	},
	{
		InputA: "tokk",
		InputB: "t*t",
		Output: false,
	},
	{
		InputA: "aasa",
		InputB: "s*a",
		Output: true,
	},
	{
		InputA: "ojjju",
		InputB: "*o",
		Output: true,
	},
	{
		InputA: "ckckkk",
		InputB: "ck*kc",
		Output: false,
	},
}

func Test_hasMatch(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := hasMatch(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
