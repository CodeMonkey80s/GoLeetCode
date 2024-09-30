package solution868

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output int
}{
	{
		Input:  22,
		Output: 2,
	},
	{
		Input:  8,
		Output: 0,
	},
	{
		Input:  5,
		Output: 2,
	},
	{
		Input:  13,
		Output: 2,
	},
}

func Test_binaryGap(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := binaryGap(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkBinaryGap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = binaryGap(testCases[0].Input)
	}
}
