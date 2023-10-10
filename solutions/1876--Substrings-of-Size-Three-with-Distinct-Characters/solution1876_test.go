package solution1876

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "xyzzaz",
		Output: 1,
	},
	{
		Input:  "aababcabc",
		Output: 4,
	},
}

func Test_countGoodSubstrings(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countGoodSubstrings(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countGoodSubstrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countGoodSubstrings(testCases[0].Input)
	}
}
