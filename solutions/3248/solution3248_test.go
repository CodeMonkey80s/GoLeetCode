package solution3248

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA int
	InputB []string
	Output int
}{
	{
		InputA: 2,
		InputB: []string{"RIGHT", "DOWN"},
		Output: 3,
	},
	{
		InputA: 3,
		InputB: []string{"DOWN", "RIGHT", "UP"},
		Output: 1,
	},
}

func Test_finalPositionOfSnake(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := finalPositionOfSnake(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_finalPositionOfSnake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = finalPositionOfSnake(testCases[0].InputA, testCases[0].InputB)
	}
}
