package solution383

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
		InputA: "a",
		InputB: "b",
		Output: false,
	},
	{
		InputA: "aa",
		InputB: "ab",
		Output: false,
	},
	{
		InputA: "aa",
		InputB: "aab",
		Output: true,
	},
	// Additional my custom cases
	{
		InputA: "aab",
		InputB: "baa",
		Output: true,
	},
}

func Test_canConstruct(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := canConstruct(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_canConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = canConstruct(testCases[0].InputA, testCases[0].InputB)
	}
}
