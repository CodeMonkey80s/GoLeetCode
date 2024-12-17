package solution3340

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	{
		Input:  "1234",
		Output: false,
	},
	{
		Input:  "24123",
		Output: true,
	},
}

func Test_isBalanced(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isBalanced(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isBalanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isBalanced(testCases[0].Input)
	}
}
