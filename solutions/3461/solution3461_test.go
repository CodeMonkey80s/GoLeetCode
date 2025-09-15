package solution3561

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	{
		Input:  "3902",
		Output: true,
	},
	{
		Input:  "34789",
		Output: false,
	},
}

func Test_hasSameDigits(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := hasSameDigits(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_makeGood(b *testing.B) {
	for b.Loop() {
		_ = hasSameDigits(testCases[0].Input)
	}
}
