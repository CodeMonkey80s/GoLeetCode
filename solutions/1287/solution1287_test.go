package solution1287

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 2, 2, 6, 6, 6, 6, 7, 10},
		Output: 6,
	},
	{
		Input:  []int{1, 1},
		Output: 1,
	},
}

func Test_findSpecialInteger(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findSpecialInteger(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findSpecialInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findSpecialInteger(testCases[0].Input)
	}
}
