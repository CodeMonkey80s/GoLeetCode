package solution2451

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output string
}{
	{
		Input:  []string{"adc", "wzy", "abc"},
		Output: "abc",
	},
	{
		Input:  []string{"aaa", "bob", "ccc", "ddd"},
		Output: "bob",
	},
}

func Test_oddString(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := oddString(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_oddString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = oddString(testCases[0].Input)
	}
}
