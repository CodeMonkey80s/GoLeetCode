package solution1670

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
			{3, 2, 1},
		},
		Output: 6,
	},
	{
		Input: [][]int{
			{1, 5},
			{7, 3},
			{3, 5},
		},
		Output: 10,
	},
	{
		Input: [][]int{
			{2, 8, 7},
			{7, 1, 3},
			{1, 9, 5},
		},
		Output: 17,
	},
}

func Test_maximumWealth(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maximumWealth(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maximumWealth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maximumWealth(testCases[0].Input)
	}
}
