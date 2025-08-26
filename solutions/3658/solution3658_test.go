package solution3658

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output int
}{

	{
		Input:  4,
		Output: 4,
	},
	{
		Input:  5,
		Output: 5,
	},
}

func Test_gcdOfOddEvenSums(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := gcdOfOddEvenSums(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_gcdOfOddEvenSums(b *testing.B) {
	for b.Loop() {
		_ = gcdOfOddEvenSums(testCases[0].Input)
	}
}
