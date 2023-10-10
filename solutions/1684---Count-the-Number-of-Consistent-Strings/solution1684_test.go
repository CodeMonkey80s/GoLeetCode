package solution1684

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB []string
	Output int
}{
	{
		InputA: "ab",
		InputB: []string{"ad", "bd", "aaab", "baa", "badab"},
		Output: 2,
	},
	{
		InputA: "abc",
		InputB: []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
		Output: 7,
	},
	{
		InputA: "cad",
		InputB: []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
		Output: 4,
	},
}

func Test_countConsistentStrings(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countConsistentStrings(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countConsistentStrings_map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countConsistentStrings(testCases[2].InputA, testCases[2].InputB)
	}
}

func Benchmark_countConsistentStrings_array(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countConsistentStrings_array(testCases[2].InputA, testCases[2].InputB)
	}
}

func Benchmark_countConsistentStrings_slice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countConsistentStrings_slice(testCases[2].InputA, testCases[2].InputB)
	}
}
