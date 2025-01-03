package solution171

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
		Input:  "A",
		Output: 1,
	},
	{
		Input:  "Z",
		Output: 26,
	},
	{
		Input:  "Y",
		Output: 25,
	},
	{
		Input:  "AB",
		Output: 28,
	},
	{
		Input:  "ZY",
		Output: 701,
	},
	// Additional my custom cases
}

func Test_titleToNumber(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := titleToNumber(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_titleToNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = titleToNumber(testCases[0].Input)
	}
}
