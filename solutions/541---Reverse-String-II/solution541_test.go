package solution541

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB int
	Output string
}{
	// Mandatory Test Cases
	{
		InputA: "abcdefg",
		InputB: 2,
		Output: "bacdfeg",
	},
	{
		InputA: "abcd",
		InputB: 2,
		Output: "bacd",
	},
	// Additional my custom cases
	{
		InputA: "abcd",
		InputB: 4,
		Output: "dcba",
	},
}

func Test_reverseStr(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reverseStr(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_reverseStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverseStr(testCases[0].InputA, testCases[0].InputB)
	}
}
