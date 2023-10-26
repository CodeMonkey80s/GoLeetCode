package solution1967

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
		InputA: []string{"a", "abc", "bc", "d"},
		InputB: "abc",
		Output: 3,
	},
	{
		InputA: []string{"a", "b", "c"},
		InputB: "aaaaabbbbb",
		Output: 2,
	},
	{
		InputA: []string{"a", "a", "a"},
		InputB: "ab",
		Output: 3,
	},
}

func Test_numOfStrings(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numOfStrings(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_numOfStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numOfStrings(testCases[0].InputA, testCases[0].InputB)
	}
}
