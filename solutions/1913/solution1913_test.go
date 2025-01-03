package solution1913

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{5, 6, 2, 7, 4},
		Output: 34,
	},
	{
		Input:  []int{4, 2, 5, 9, 7, 4, 8},
		Output: 64,
	},
}

func Test_maxProductDifference(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maxProductDifference(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maxProductDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxProductDifference(testCases[0].Input)
	}
}
