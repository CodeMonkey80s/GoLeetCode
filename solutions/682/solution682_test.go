package solution682

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input:  []string{"5", "2", "C", "D", "+"},
		Output: 30,
	},
	{
		Input:  []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
		Output: 27,
	},
	{
		Input:  []string{"1", "C"},
		Output: 0,
	},
}

func Test_calPoints(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := calPoints(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_calPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = calPoints(testCases[0].Input)
	}
}
