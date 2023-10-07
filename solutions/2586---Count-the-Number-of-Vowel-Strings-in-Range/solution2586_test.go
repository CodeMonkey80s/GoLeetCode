package solution2586

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB int
	InputC int
	Output int
}{
	{
		InputA: []string{"are", "amy", "u"},
		InputB: 0,
		InputC: 2,
		Output: 2,
	},
	{
		InputA: []string{"hey", "aeo", "mu", "ooo", "artro"},
		InputB: 1,
		InputC: 4,
		Output: 3,
	},
}

func Test_vowelStrings(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v %v Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := vowelStrings(tc.InputA, tc.InputB, tc.InputC)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_vowelStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = vowelStrings(testCases[0].InputA, testCases[0].InputB, testCases[0].InputC)
	}
}
