package solution2185

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB string
	Output int
}{
	{
		InputA: []string{"pay", "attention", "practice", "attend"},
		InputB: "at",
		Output: 2,
	},
	{
		InputA: []string{"leetcode", "win", "loops", "success"},
		InputB: "code",
		Output: 0,
	},
}

func Test_prefixCount(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := prefixCount(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_prefixCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prefixCount(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_prefixCount_strings_index(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prefixCount_strings_index(testCases[0].InputA, testCases[0].InputB)
	}
}
