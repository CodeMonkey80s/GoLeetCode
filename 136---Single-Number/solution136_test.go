package solution136

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  []int{2, 2, 1},
		Output: 1,
	},
	{
		Input:  []int{4, 1, 2, 1, 2},
		Output: 4,
	},
	{
		Input:  []int{1},
		Output: 1,
	},
	// Additional my custom cases
}

func Test_singleNumbert(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := singleNumber(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_singleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = singleNumber(testCases[0].Input)
	}
}
