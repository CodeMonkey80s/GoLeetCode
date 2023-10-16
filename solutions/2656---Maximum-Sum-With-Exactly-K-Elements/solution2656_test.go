package solution2656

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
		InputA: []int{1, 2, 3, 4, 5},
		InputB: 3,
		Output: 18,
	},
	{
		InputA: []int{5, 5, 5},
		InputB: 2,
		Output: 11,
	},
}

func Test_maximizeSum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maximizeSum(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maximizeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maximizeSum(testCases[0].InputA, testCases[0].InputB)
	}
}
