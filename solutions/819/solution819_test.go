package solution819

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB []string
	Output string
}{
	// Mandatory Test Cases
	{
		InputA: "Bob hit a ball, the hit BALL flew far after it was hit.",
		InputB: []string{"hit"},
		Output: "ball",
	},
	{
		InputA: "a, a, a, a, b,b,b,c, c",
		InputB: []string{"a"},
		Output: "b",
	},
	{
		InputA: "a",
		InputB: []string{""},
		Output: "a",
	},
	// Additional my custom cases
}

func Test_mostCommonWord(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := mostCommonWord(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_mostCommonWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mostCommonWord(testCases[0].InputA, testCases[0].InputB)
	}
}
