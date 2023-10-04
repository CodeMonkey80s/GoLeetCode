package solution557

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
		Input:  "Let's take LeetCode contest",
		Output: "s'teL ekat edoCteeL tsetnoc",
	},
	{
		Input:  "God Ding",
		Output: "doG gniD",
	},
	// Additional my custom cases
}

func Test_reverseWords(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reverseWords(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_reverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverseWords(testCases[0].Input)
	}
}
