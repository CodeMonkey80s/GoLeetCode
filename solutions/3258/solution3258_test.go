package solution3258

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB int
	Output int
}{
	{
		InputA: "10101",
		InputB: 1,
		Output: 12,
	},
	{
		InputA: "1010101",
		InputB: 2,
		Output: 25,
	},
	{
		InputA: "11111",
		InputB: 1,
		Output: 15,
	},
	{
		InputA: "000011",
		InputB: 1,
		Output: 18,
	},
}

func Test_countKConstraintSubstrings(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countKConstraintSubstrings(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countKConstraintSubstrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countKConstraintSubstrings(testCases[0].InputA, testCases[0].InputB)
	}
}
