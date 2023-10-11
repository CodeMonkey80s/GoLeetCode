package solution1844

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "a1c1e1",
		Output: "abcdef",
	},
	{
		Input:  "a1b2c3d4e",
		Output: "abbdcfdhe",
	},
}

func Test_replaceDigits(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := replaceDigits(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func Benchmark_replaceDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = replaceDigits(testCases[0].Input)
	}
}
