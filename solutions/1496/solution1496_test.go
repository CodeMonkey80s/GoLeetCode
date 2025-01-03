package solution1496

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
		Input:  "NES",
		Output: false,
	},
	{
		Input:  "NESWW",
		Output: true,
	},
	// Additional my custom cases
	{
		Input:  "NNSWWEWSSESSWENNW",
		Output: true,
	},
}

func Test_isPathCrossing(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isPathCrossing(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isPathCrossing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPathCrossing(testCases[0].Input)
	}
}
