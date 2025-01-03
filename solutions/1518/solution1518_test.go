package solution1518

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
		InputA: 9,
		InputB: 3,
		Output: 13,
	},
	{
		InputA: 15,
		InputB: 4,
		Output: 19,
	},
}

func Test_numWaterBottles(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numWaterBottles(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_numWaterBottles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numWaterBottles(testCases[0].InputA, testCases[0].InputB)
	}
}
