package solution520

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
		Input:  "USA",
		Output: true,
	},
	{
		Input:  "leetcode",
		Output: true,
	},
	{
		Input:  "Google",
		Output: true,
	},
	{
		Input:  "FlaG",
		Output: false,
	},
	// Additional my custom cases
	{
		Input:  "FFFFFFFFFFFFFFFFFFFFf",
		Output: false,
	},
	{
		Input:  "AMAzON",
		Output: false,
	},
}

func Test_detectCapitalUse(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := detectCapitalUse_first_attempt(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_detectCapitalUse_std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = detectCapitalUse(testCases[0].Input)
	}
}

func Benchmark_detectCapitalUse_first_attempt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = detectCapitalUse_first_attempt(testCases[0].Input)
	}
}
