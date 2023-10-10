package solution2283

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	{
		Input:  "1210",
		Output: true,
	},
	{
		Input:  "030",
		Output: false,
	},
}

func Test_digitCount(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := digitCount(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_digitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = digitCount(testCases[0].Input)
	}
}
