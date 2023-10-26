package solution2423

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	{
		Input:  "abcc",
		Output: true,
	},
	{
		Input:  "aazz",
		Output: false,
	},
	{
		Input:  "adbc",
		Output: true,
	},
	{
		Input:  "ddaccb",
		Output: false,
	},
	{
		Input:  "cccaa",
		Output: true,
	},
}

func Test_equalFrequency(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := equalFrequency(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_equalFequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = equalFrequency(testCases[0].Input)
	}
}
