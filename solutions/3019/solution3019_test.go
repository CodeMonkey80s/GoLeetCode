package solution3019

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "aAbBcC",
		Output: 2,
	},
	{
		Input:  "AaAaAaaA",
		Output: 0,
	},
}

func Test_countKeyChanges(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countKeyChanges(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countKeyChanges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countKeyChanges(testCases[0].Input)
	}
}
