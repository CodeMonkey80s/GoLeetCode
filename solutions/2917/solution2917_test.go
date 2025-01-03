package solution2917

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
		InputA: []int{7, 12, 9, 8, 9, 15},
		InputB: 4,
		Output: 9,
	},
	{
		InputA: []int{2, 12, 1, 11, 4, 5},
		InputB: 6,
		Output: 0,
	},
	{
		InputA: []int{10, 8, 5, 9, 11, 6, 8},
		InputB: 1,
		Output: 15,
	},
}

func Test_findKOr(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findKOrV2(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findKOrV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findKOrV2(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_findKOrV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findKOrV1(testCases[0].InputA, testCases[0].InputB)
	}
}
