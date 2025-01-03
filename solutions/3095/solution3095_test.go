package solution3095

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
		InputA: []int{1, 2, 3},
		InputB: 2,
		Output: 1,
	},
	{
		InputA: []int{2, 1, 8},
		InputB: 10,
		Output: 3,
	},
	{
		InputA: []int{1, 2},
		InputB: 0,
		Output: 1,
	},
	{
		InputA: []int{1, 12, 2, 5},
		InputB: 43,
		Output: -1,
	},
}

func Test_minimumSubarrayLength(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minimumSubarrayLength(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minimumSubarrayLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minimumSubarrayLength(testCases[0].InputA, testCases[0].InputB)
	}
}
