package solution20

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	// Mandatory Test Cases
	{
		Input:  "()",
		Output: true,
	},
	{
		Input:  "()[]{}",
		Output: true,
	},
	{
		Input:  "(])",
		Output: false,
	},
	// Additional my custom cases
	{
		Input:  "([)]",
		Output: false,
	},
	{
		Input:  "{[]}",
		Output: true,
	},
	{
		Input:  ")(){}",
		Output: false,
	},
	{
		Input:  "[])",
		Output: false,
	},
	{
		Input:  "()))",
		Output: false,
	},
	{
		Input:  "([}}])",
		Output: false,
	},
}

func Test_isValid(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isValid(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isValid(testCases[0].Input)
	}
}
