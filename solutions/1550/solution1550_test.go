package solution1550

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output bool
}{
	{
		Input:  []int{2, 6, 4, 1},
		Output: false,
	},
	{
		Input:  []int{1, 2, 34, 3, 4, 5, 7, 23, 12},
		Output: true,
	},
}

func Test_threeConsecutiveOdds(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := threeConsecutiveOdds(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_threeConsecutiveOdds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = threeConsecutiveOdds(testCases[0].Input)
	}
}
