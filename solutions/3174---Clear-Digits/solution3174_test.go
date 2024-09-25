package solution3174

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "abc",
		Output: "abc",
	},
	{
		Input:  "cb34",
		Output: "",
	},
	{
		Input:  "a",
		Output: "a",
	},
	{
		Input:  "b4y6",
		Output: "",
	},
}

func Test_clearDigits(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %q, Output: %q\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := clearDigits(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func Benchmark_ClearDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = clearDigits(testCases[0].Input)
	}
}
