package solution459

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	// Mandatory Test Cases
	{
		Input:  "abab",
		Output: true,
	},
	{
		Input:  "aba",
		Output: false,
	},
	{
		Input:  "abcabcabcabc",
		Output: true,
	},
	// Additional my custom cases
}

func Test_repeatedSubstringPatterns(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := repeatedSubstringPattern(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_repeatedSubstringPatterns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = repeatedSubstringPattern(testCases[0].Input)
	}
}
