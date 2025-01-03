package solution2824

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
		InputA: []int{-1, 1, 2, 3, 1},
		InputB: 2,
		Output: 3,
	},
	{
		InputA: []int{-6, 2, 5, -2, -7, -1, 3},
		InputB: -2,
		Output: 10,
	},
}

func Test_countPairs(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countPairs(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countPairs(testCases[0].InputA, testCases[0].InputB)
	}
}
