package solution575

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 1, 2, 2, 3, 3},
		Output: 3,
	},
	{
		Input:  []int{1, 1, 2, 3},
		Output: 2,
	},
	{
		Input:  []int{6, 6, 6, 6},
		Output: 1,
	},
}

func Test_distributeCandies(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := distributeCandies(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_distributeCandies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = distributeCandies(testCases[0].Input)
	}
}
