package solution3151

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output bool
}{
	{
		Input:  []int{1},
		Output: true,
	},
	{
		Input:  []int{2, 1, 4},
		Output: true,
	},
	{
		Input:  []int{4, 3, 1, 6},
		Output: false,
	},
}

func Test_isArraySpecial(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isArraySpecial(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isArraySpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isArraySpecial(testCases[2].Input)
	}
}
