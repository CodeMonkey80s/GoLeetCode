package solution3074

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	Output int
}{
	{
		InputA: []int{1, 3, 2},
		InputB: []int{4, 3, 1, 5, 2},
		Output: 2,
	},
	{
		InputA: []int{5, 5, 5},
		InputB: []int{2, 4, 2, 7},
		Output: 4,
	},
}

func Test_minimumBoxes(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minimumBoxes(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minimumBoxes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minimumBoxes(testCases[0].InputA, testCases[0].InputB)
	}
}
