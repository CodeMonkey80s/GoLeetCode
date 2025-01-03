package solution2716

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "aaabc",
		Output: 3,
	},
	{
		Input:  "cbbd",
		Output: 3,
	},
	{
		Input:  "dddaaa",
		Output: 2,
	},
}

func Test_minimizedStringLength(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minimizedStringLength(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minimizeStringLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minimizedStringLength(testCases[0].Input)
	}
}
