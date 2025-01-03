package solution345

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	// Mandatory Test Cases
	{
		Input:  "hello",
		Output: "holle",
	},
	{
		Input:  "leetcode",
		Output: "leotcede",
	},
	// Additional my custom cases
	{
		Input:  `Marge, let's "went." I await news telegram.`,
		Output: `Marge, let's "went." i awaIt news telegram.`,
	},
}

func Test_reverseVowels(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reverseVowels(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_reverseVowels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverseVowels(testCases[0].Input)
	}
}
