package solution1528

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB []int
	Output string
}{
	// Mandatory Test Cases
	{
		InputA: "codeleet",
		InputB: []int{4, 5, 6, 7, 0, 2, 1, 3},
		Output: "leetcode",
	},
	{
		InputA: "abc",
		InputB: []int{0, 1, 2},
		Output: "abc",
	},
	// Additional my custom cases
}

func Test_restoreString(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := restoreString(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_restoreString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = restoreString(testCases[0].InputA, testCases[0].InputB)
	}
}
