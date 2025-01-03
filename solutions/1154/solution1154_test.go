package solution1154

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  "2019-01-09",
		Output: 9,
	},
	{
		Input:  "2019-02-10",
		Output: 41,
	},
	// Additional my custom cases
}

func Test_dayOfYear(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := dayOfYear(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_dayOfYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = dayOfYear(testCases[0].Input)
	}
}
