package solution3550

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 3, 2},
		Output: 2,
	},
	{
		Input:  []int{1, 10, 11},
		Output: 1,
	},
	{
		Input:  []int{1, 2, 3},
		Output: -1,
	},
}

func Test_smallestIndex(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := smallestIndex(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_smallestInex(b *testing.B) {
	for b.Loop() {
		_ = smallestIndex(testCases[0].Input)
	}
}
