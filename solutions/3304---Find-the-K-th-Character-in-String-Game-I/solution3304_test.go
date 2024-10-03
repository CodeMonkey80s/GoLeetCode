package solution3304

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output byte
}{
	{
		Input:  5,
		Output: 'b',
	},
	{
		Input:  10,
		Output: 'c',
	},
}

func Test_kthCharacter(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := kthCharacter(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_kthCharacter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = kthCharacter(testCases[0].Input)
	}
}
