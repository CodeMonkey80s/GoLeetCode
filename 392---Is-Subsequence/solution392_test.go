package solution392

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output bool
}{
	// Mandatory Test Cases
	{
		InputA: "abc",
		InputB: "ahbgdc",
		Output: true,
	},
	{
		InputA: "axc",
		InputB: "ahbgdc",
		Output: false,
	},
	// Additional my custom cases
	{
		InputA: "acb",
		InputB: "ahbgdc",
		Output: false,
	},
	{
		InputA: "aaaaaa",
		InputB: "bbaaaa",
		Output: false,
	},
	{
		InputA: "bb",
		InputB: "ahbgdc",
		Output: false,
	},
}

func Test_isSubsequence(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isSubsequence(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isSubsequence(testCases[0].InputA, testCases[0].InputB)
	}
}
