package solution215

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
		InputA: []int{3, 2, 1, 5, 6, 4},
		InputB: 2,
		Output: 5,
	},
	{
		InputA: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		InputB: 4,
		Output: 4,
	},
	{
		InputA: []int{2, 1},
		InputB: 2,
		Output: 1,
	},
}

func Test_findKthLargest(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findKthLargestV2(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findKthLargestV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findKthLargestV2(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_findKthLargestV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findKthLargestV1(testCases[0].InputA, testCases[0].InputB)
	}
}
