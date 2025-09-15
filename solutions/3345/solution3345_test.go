package solution3345

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA int
	InputB int
	Output int
}{
	{
		InputA: 10,
		InputB: 2,
		Output: 10,
	},
	{
		InputA: 15,
		InputB: 3,
		Output: 16,
	},
}

func Test_smallestNumber(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := smallestNumber(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_smallestNumber(b *testing.B) {
	for b.Loop() {
		_ = smallestNumber(testCases[0].InputA, testCases[0].InputB)
	}
}
