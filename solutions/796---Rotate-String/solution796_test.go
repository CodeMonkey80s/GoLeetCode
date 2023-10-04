package solution796

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
		InputA: "abcde",
		InputB: "cdeab",
		Output: true,
	},
	{
		InputA: "abcde",
		InputB: "abced",
		Output: false,
	},
	{
		InputA: "abcde",
		InputB: "abcde",
		Output: true,
	},
	// Additional my custom cases
}

func Test_rotateString(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := rotateString_copy(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_rotateString_append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rotateString_append(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_rotateString_copy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rotateString_copy(testCases[0].InputA, testCases[0].InputB)
	}
}
