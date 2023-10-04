package solution771

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output int
}{
	// Mandatory Test Cases
	{
		InputA: "aA",
		InputB: "aAAbbbb",
		Output: 3,
	},
	{
		InputA: "z",
		InputB: "ZZ",
		Output: 0,
	},
	// Additional my custom cases
}

func Test_numJewles(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numJewelsInStones(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_numJewelsInStones_bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numJewelsInStones_bytes(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_numJewelsInStones_strings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numJewelsInStones_strings(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_numJewelsInStones_my(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numJewelsInStones(testCases[0].InputA, testCases[0].InputB)
	}
}
