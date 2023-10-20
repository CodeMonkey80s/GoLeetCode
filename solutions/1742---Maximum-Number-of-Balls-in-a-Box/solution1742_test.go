package solution1742

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
		InputB: 10,
		Output: 2,
	},
	{
		InputA: 5,
		InputB: 15,
		Output: 2,
	},
	{
		InputA: 19,
		InputB: 28,
		Output: 2,
	},
}

func Test_countBalls(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countBalls(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countBalls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countBalls(testCases[0].InputA, testCases[0].InputB)
	}
}
