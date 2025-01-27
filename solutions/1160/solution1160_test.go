package solution1160

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
		InputA: []string{"cat", "bt", "hat", "tree"},
		InputB: "atach",
		Output: 6,
	},
	{
		InputA: []string{"hello", "world", "leetcode"},
		InputB: "welldonehoneyr",
		Output: 10,
	},
}

func Test_countCharacters(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countCharacters(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countCharacters(testCases[0].InputA, testCases[0].InputB)
	}
}
