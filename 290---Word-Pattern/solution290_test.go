package solution290

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
		InputA: "abba",
		InputB: "dog cat cat dog",
		Output: true,
	},
	{
		InputA: "abba",
		InputB: "dog cat cat fish",
		Output: false,
	},
	{
		InputA: "aaaa",
		InputB: "dog cat cat dog",
		Output: false,
	},
	// Additional my custom cases
	{
		InputA: "abba",
		InputB: "dog dog dog dog",
		Output: false,
	},
	{
		InputA: "aaa",
		InputB: "aa aa aa aa",
		Output: false,
	},
}

func Test_wordPattern(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := wordPattern(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_wordPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = wordPattern(testCases[0].InputA, testCases[0].InputB)
	}
}
