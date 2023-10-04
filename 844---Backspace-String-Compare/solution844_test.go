package solution844

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output bool
}{
	// Mandatory Test Cases
	{
		InputA: "ab#c",
		InputB: "ad#c",
		Output: true,
	},
	{
		InputA: "ab##",
		InputB: "c#d#",
		Output: true,
	},
	{
		InputA: "a#c",
		InputB: "b",
		Output: false,
	},
	// Additional my custom cases
	{
		InputA: "a##c",
		InputB: "#a#c",
		Output: true,
	},
}

func Test_backspaceCompare(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := backspaceCompare(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_backspaceCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = backspaceCompare(testCases[0].InputA, testCases[0].InputB)
	}
}
