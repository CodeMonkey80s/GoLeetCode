package solution1768

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output string
}{
	{
		InputA: "abc",
		InputB: "pqr",
		Output: "apbqcr",
	},
	{
		InputA: "ab",
		InputB: "pqrs",
		Output: "apbqrs",
	},
	{
		InputA: "abcd",
		InputB: "pq",
		Output: "apbqcd",
	},
}

func Test_mergeAlternately(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := mergeAlternately(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_mergeAlternately(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mergeAlternately(testCases[0].InputA, testCases[0].InputB)
	}
}
