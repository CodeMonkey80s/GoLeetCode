package solution58

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
		Input:  "Hello World",
		Output: 5,
	},
	{
		Input:  "   fly me   to   the moon  ",
		Output: 4,
	},
	{
		Input:  "luffy is still joyboy",
		Output: 6,
	},
	// Additional my custom cases
}

func Test_lengthOfLastWords(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := lengthOfLastWord(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_lengthOfLongestSubstring_stringsField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lengthOfLastWord_stringsField(testCases[0].Input)
	}
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lengthOfLastWord(testCases[0].Input)
	}
}
