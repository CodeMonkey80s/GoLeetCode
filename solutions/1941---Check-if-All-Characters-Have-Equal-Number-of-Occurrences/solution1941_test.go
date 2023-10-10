package solution1941

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	{
		Input:  "abacbc",
		Output: true,
	},
	{
		Input:  "aaabb",
		Output: false,
	},
	{
		Input:  "vvvvvvvvvvvvvvvvvvv",
		Output: true,
	},
}

func Test_areOccurrencesEqual(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := areOccurrencesEqual(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_areOccurrencesEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = areOccurrencesEqual(testCases[0].Input)
	}
}

func Benchmark_areOccurrencesEqual_map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = areOccurrencesEqual_map(testCases[0].Input)
	}
}
