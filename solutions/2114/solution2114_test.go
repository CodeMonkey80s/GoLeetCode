package solution2114

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input:  []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"},
		Output: 6,
	},
	{
		Input:  []string{"please wait", "continue to fight", "continue to win"},
		Output: 3,
	},
}

func Test_mostWordsFound(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := mostWordsFound(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_mostWordsFound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mostWordsFound(testCases[0].Input)
	}
}

func Benchmark_mostWordsFound_strings_count(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mostWordsFound_strings_count(testCases[0].Input)
	}
}
