package solution1816

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB int
	Output string
}{
	{
		InputA: "Hello how are you Contestant",
		InputB: 4,
		Output: "Hello how are you",
	},
	{
		InputA: "What is the solution to this problem",
		InputB: 4,
		Output: "What is the solution",
	},
	{
		InputA: "chopper is not a tanuki",
		InputB: 5,
		Output: "chopper is not a tanuki",
	},
}

func Test_truncateSentence(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := truncateSentence(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func Benchmark_truncateSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = truncateSentence(testCases[0].InputA, testCases[0].InputB)
	}
}
