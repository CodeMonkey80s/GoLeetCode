package solution1588

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 4, 2, 5, 3},
		Output: 58,
	},
	{
		Input:  []int{1, 2},
		Output: 3,
	},
	{
		Input:  []int{10, 11, 12},
		Output: 66,
	},
}

func Test_sumOddLengthsSubarrays(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sumOddLengthSubarrays(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_sumOddLengthsSubarrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sumOddLengthSubarrays(testCases[0].Input)
	}
}
