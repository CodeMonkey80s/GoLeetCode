package solution3

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  "abcabcbb",
		Output: 3,
	},
	{
		Input:  "bbbbb",
		Output: 1,
	},
	{
		Input:  "pwwkew",
		Output: 3,
	},
	// Additional my custom cases
	{
		Input:  "dvdf",
		Output: 3,
	},
	{
		Input:  "cdd",
		Output: 2,
	},
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := lengthOfLongestSubstring(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lengthOfLongestSubstring(testCases[0].Input)
	}
}
