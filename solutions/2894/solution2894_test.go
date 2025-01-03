package solution2894

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
		InputB: 3,
		Output: 19,
	},
}

func Test_differenceOfSums(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := differenceOfSums(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_differenceOfSums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = differenceOfSums(testCases[0].InputA, testCases[0].InputB)
	}
}
