package solution3046

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output bool
}{
	{
		Input:  []int{1, 1, 2, 2, 3, 4},
		Output: true,
	},
	{
		Input:  []int{1, 1, 1, 1},
		Output: false,
	},
}

func Test_isPossibleToSplit(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isPossibleToSplit(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isPossibleToSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPossibleToSplit(testCases[0].Input)
	}
}
