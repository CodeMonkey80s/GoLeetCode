package solution1051

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 1, 4, 2, 1, 3},
		Output: 3,
	},
	{
		Input:  []int{5, 1, 2, 3, 4},
		Output: 5,
	},
	{
		Input:  []int{1, 2, 3, 4, 5},
		Output: 0,
	},
}

func Test_heightChecker(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := heightChecker(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_heightChecker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = heightChecker(testCases[0].Input)
	}
}
