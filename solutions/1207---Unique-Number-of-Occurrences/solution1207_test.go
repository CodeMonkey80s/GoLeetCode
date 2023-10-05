package solution1207

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output bool
}{
	{
		Input:  []int{1, 2, 2, 1, 1, 3},
		Output: true,
	},
	{
		Input:  []int{1, 2},
		Output: false,
	},
	{
		Input:  []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
		Output: true,
	},
}

func Test_unique(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := uniqueOccurrences(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_unique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uniqueOccurrences(testCases[0].Input)
	}
}
