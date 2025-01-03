package solution3226

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
		InputA: 13,
		InputB: 4,
		Output: 2,
	},
	{
		InputA: 21,
		InputB: 21,
		Output: 0,
	},
	{
		InputA: 14,
		InputB: 13,
		Output: -1,
	},
}

func Test_minChanges(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minChanges(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minChanges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minChanges(testCases[0].InputA, testCases[0].InputB)
	}
}
