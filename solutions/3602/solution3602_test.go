package solution3602

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output string
}{
	{
		Input:  13,
		Output: "A91P1",
	},
	{
		Input:  36,
		Output: "5101000",
	},
}

func Test_concatHex36(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := concatHex36(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_concatHex36(b *testing.B) {
	for b.Loop() {
		_ = concatHex36(testCases[0].Input)
	}
}
