package solution3370

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output int
}{
	{
		Input:  5,
		Output: 7,
	},
	{
		Input:  10,
		Output: 15,
	},
	{
		Input:  3,
		Output: 3,
	},
}

func Test_smallestNumber(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := smallestNumber(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isBalanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = smallestNumber(testCases[0].Input)
	}
}
