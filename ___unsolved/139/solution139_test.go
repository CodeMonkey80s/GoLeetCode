package solution139

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB []string
	Output bool
}{
	{
		InputA: "leetcode",
		InputB: []string{"leet", "code"},
		Output: true,
	},
	{
		InputA: "applepenapple",
		InputB: []string{"apple", "pen"},
		Output: true,
	},
	{
		InputA: "catsandog",
		InputB: []string{"cats", "dog", "sand", "and", "cat"},
		Output: false,
	},
	{
		InputA: "bb",
		InputB: []string{"a", "b", "bbb", "bbbb"},
		Output: true,
	},
}

func Test_wordBreak(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := wordBreak(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
