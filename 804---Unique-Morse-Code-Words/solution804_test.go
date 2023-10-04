package solution804

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  []string{"gin", "zen", "gig", "msg"},
		Output: 2,
	},
	{
		Input:  []string{"a"},
		Output: 1,
	},
	// Additional my custom cases
}

func Test_uniqueMorse(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := uniqueMorseRepresentations(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_uniqueMorse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uniqueMorseRepresentations(testCases[0].Input)
	}
}
