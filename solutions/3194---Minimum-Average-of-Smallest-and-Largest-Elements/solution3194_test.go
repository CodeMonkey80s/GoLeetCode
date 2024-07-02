package solution3194

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output float64
}{
	{
		Input:  []int{7, 8, 3, 4, 15, 13, 4, 1},
		Output: 5.5,
	},
	{
		Input:  []int{1, 9, 8, 3, 10, 5},
		Output: 5.5,
	},
	{
		Input:  []int{1, 2, 3, 7, 8, 9},
		Output: 5.0,
	},
}

func Test_minimumAverage(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minimumAverage(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minimumAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minimumAverage(testCases[0].Input)
	}
}
