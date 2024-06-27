package solution2843

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
		InputA: 1,
		InputB: 100,
		Output: 9,
	},
	{
		InputA: 1200,
		InputB: 1230,
		Output: 4,
	},
}

func Test_countSymmetricIntegers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countSymmetricIntegers(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countSymmetricIntegers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countSymmetricIntegers(testCases[0].InputA, testCases[0].InputB)
	}
}
