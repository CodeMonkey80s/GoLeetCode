package solution3541

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{

	{
		Input:  "successes",
		Output: 6,
	},
	{
		Input:  "aeiaeia",
		Output: 3,
	},
}

func Test_maxFreqSum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maxFreqSum(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maxFreqSum(b *testing.B) {
	for b.Loop() {
		_ = maxFreqSum(testCases[0].Input)
	}
}
