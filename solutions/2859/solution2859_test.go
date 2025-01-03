package solution2859

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output int
}{
	{
		InputA: []int{5, 10, 1, 5, 2},
		InputB: 1,
		Output: 13,
	},
	{
		InputA: []int{4, 3, 2, 1},
		InputB: 2,
		Output: 1,
	},
}

func Test_sumIndices(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sumIndicesWithKSetBits(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_sumIndices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sumIndicesWithKSetBits(testCases[0].InputA, testCases[0].InputB)
	}
}
