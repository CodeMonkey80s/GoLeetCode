package solution3158

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 2, 1, 3},
		Output: 1,
	},
	{
		Input:  []int{1, 2, 3},
		Output: 0,
	},
	{
		Input:  []int{1, 2, 2, 1},
		Output: 3,
	},
	{
		Input:  []int{1, 4, 4},
		Output: 4,
	},
	{
		Input:  []int{1, 2, 3, 3, 2, 1},
		Output: 0,
	},
	{
		Input:  []int{1, 2, 10, 4, 2, 5, 4, 5, 8, 10},
		Output: 9,
	},
}

func Test_duplicateNumbersXOR(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := duplicateNumbersXOR(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_duplicateNumbersXOR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = duplicateNumbersXOR(testCases[4].Input)
	}
}
