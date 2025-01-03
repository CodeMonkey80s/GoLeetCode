package solution762

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
		InputA: 6,
		InputB: 10,
		Output: 4,
	},
	{
		InputA: 10,
		InputB: 15,
		Output: 5,
	},
}

func Test_countPrimeSetBits(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countPrimeSetBits(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countPrimeSetBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countPrimeSetBits(testCases[0].InputA, testCases[0].InputB)
	}
}
