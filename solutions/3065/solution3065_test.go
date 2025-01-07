package solution3065

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
		InputA: []int{2, 11, 10, 1, 3},
		InputB: 10,
		Output: 3,
	},
	{
		InputA: []int{1, 1, 2, 4, 9},
		InputB: 1,
		Output: 0,
	},
	{
		InputA: []int{1, 1, 2, 4, 9},
		InputB: 9,
		Output: 4,
	},
}

func Test_minOperations(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minOperations(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minOperations(testCases[0].InputA, testCases[0].InputB)
	}
}
