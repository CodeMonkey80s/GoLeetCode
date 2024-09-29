package solution963

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output bool
}{
	{
		Input:  5,
		Output: true,
	},
	{
		Input:  7,
		Output: false,
	},
	{
		Input:  11,
		Output: false,
	},
}

func Test_hasAlternatingBits(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := hasAlternatingBits(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_hasAlternatingBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hasAlternatingBits(testCases[0].Input)
	}
}
