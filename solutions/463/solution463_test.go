package solution463

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output int
}{
	{
		Input:  [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
		Output: 16,
	},
	{
		Input:  [][]int{{1}},
		Output: 4,
	},
	{
		Input:  [][]int{{1, 0}},
		Output: 4,
	},
	{
		Input:  [][]int{{0, 1}},
		Output: 4,
	},
	{
		Input:  [][]int{{1, 1}},
		Output: 6,
	},
}

func Test_islandPerimeter(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := islandPerimeter(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_islandPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = islandPerimeter(testCases[0].Input)
	}
}
