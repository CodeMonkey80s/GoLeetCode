package solution3471

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
		InputA: []int{3, 9, 2, 1, 7},
		InputB: 3,
		Output: 7,
	},
	{
		InputA: []int{3, 9, 7, 2, 1, 7},
		InputB: 4,
		Output: 3,
	},
	{
		InputA: []int{0, 0},
		InputB: 1,
		Output: -1,
	},
}

func Test_maxFreqSum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := largestInteger(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maxFreqSum(b *testing.B) {
	for b.Loop() {
		_ = largestInteger(testCases[0].InputA, testCases[0].InputB)
	}
}
