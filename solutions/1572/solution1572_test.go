package solution1572

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
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		Output: 25,
	},
	{
		Input: [][]int{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
		Output: 8,
	},
	{
		Input: [][]int{
			{5},
		},
		Output: 5,
	},
}

func Test_diagonalSum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := diagonalSum(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_diagonalSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = diagonalSum(testCases[0].Input)
	}
}
