package solution1351

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output int
}{
	{
		Input: [][]int{
			{4, 3, 2, -1},
			{3, 2, 1, -1},
			{1, 1, -1, -2},
			{-1, -1, -2, -3},
		},
		Output: 8,
	},
	{
		Input: [][]int{
			{3, 2},
			{1, 0},
		},
		Output: 0,
	},
}

func Test_countNegatives(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countNegatives(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countNegatives(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countNegatives(testCases[0].Input)
	}
}
