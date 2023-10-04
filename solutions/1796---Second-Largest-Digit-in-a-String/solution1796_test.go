package solution1796

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
		Input:  "dfa12321afd",
		Output: 2,
	},
	{
		Input:  "abc1111",
		Output: -1,
	},
	// Additional my custom cases
	{
		Input:  "ck077",
		Output: 0,
	},
}

func Test_secondHighest(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := secondHighest(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_secondHighest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = secondHighest(testCases[0].Input)
	}
}
