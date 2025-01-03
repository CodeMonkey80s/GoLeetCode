package solution2103

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "B0B6G0R6R0R6G9",
		Output: 1,
	},
	{
		Input:  "B0R0G0R9R0B0G0",
		Output: 1,
	},
	{
		Input:  "G4",
		Output: 0,
	},
}

func Test_countPoints(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countPoints(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countPoints(testCases[0].Input)
	}
}
