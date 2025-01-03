package solution205

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
		InputA: "egg",
		InputB: "add",
		Output: true,
	},
	{
		InputA: "foo",
		InputB: "bar",
		Output: false,
	},
	{
		InputA: "paper",
		InputB: "title",
		Output: true,
	},
	// Additional my custom cases
	{
		InputA: "badc",
		InputB: "baba",
		Output: false,
	},
	{
		InputA: "gg",
		InputB: "add",
		Output: false,
	},
	{
		InputA: "egcd",
		InputB: "adfd",
		Output: false,
	},
}

func Test_isIsomorphic(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isIsomorphic(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isIsomorphic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isIsomorphic(testCases[0].InputA, testCases[0].InputB)
	}
}
