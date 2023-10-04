package solution1704

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
		Input:  "book",
		Output: true,
	},
	{
		Input:  "textbook",
		Output: false,
	},
	// Additional my custom cases
}

func Test_halvesAreAlike(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := halvesAreAlike(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_halvesAreAlike(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = halvesAreAlike(testCases[0].Input)
	}
}
