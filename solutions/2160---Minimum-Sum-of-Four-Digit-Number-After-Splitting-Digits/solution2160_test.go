package solution2160

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output int
}{
	{
		Input:  2932,
		Output: 52,
	},
	{
		Input:  4009,
		Output: 13,
	},
}

func Test_minimumSum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minimumSum(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minimumSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minimumSum(testCases[0].Input)
	}
}
