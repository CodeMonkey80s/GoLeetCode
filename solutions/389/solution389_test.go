package solution389

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output byte
}{
	// Mandatory Test Cases
	{
		InputA: "abcd",
		InputB: "abcde",
		Output: byte('e'),
	},
	{
		InputA: "",
		InputB: "y",
		Output: byte('y'),
	},
	// Additional my custom cases
}

func Test_findTheDifference(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findTheDifference(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findTheDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findTheDifference(testCases[0].InputA, testCases[0].InputB)
	}
}
