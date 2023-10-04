package solution1592

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	// Mandatory Test Cases
	{
		Input:  "  this   is  a sentence ",
		Output: "this   is   a   sentence",
	},
	{
		Input:  " practice   makes   perfect",
		Output: "practice   makes   perfect ",
	},
	// Additional my custom cases
	{
		Input:  "a",
		Output: "a",
	},
	{
		Input:  "  hello",
		Output: "hello  ",
	},
}

func Test_reorderSpaces(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reorderSpaces(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_reorderSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reorderSpaces(testCases[0].Input)
	}
}
