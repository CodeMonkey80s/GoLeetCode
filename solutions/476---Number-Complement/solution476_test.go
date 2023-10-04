package solution476

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
		Output: 2,
	},
	{
		Input:  1,
		Output: 0,
	},
}

func Test_findComplement(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findComplement(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findComplement(testCases[0].Input)
	}
}
