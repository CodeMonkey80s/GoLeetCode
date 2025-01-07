package solution1561

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{

	{
		Input:  []int{2, 4, 1, 2, 7, 8},
		Output: 9,
	},
	{
		Input:  []int{2, 4, 5},
		Output: 4,
	},
	{
		Input:  []int{9, 8, 7, 6, 5, 1, 2, 3, 4},
		Output: 18,
	},
}

func Test_maxCoins(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maxCoins(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maxCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxCoins(testCases[0].Input)
	}
}
