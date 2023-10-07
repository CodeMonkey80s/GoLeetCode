package solution2696

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "ABFCACDB",
		Output: 2,
	},
	{
		Input:  "ACBBD",
		Output: 5,
	},
}

func Test_minLength(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minLength(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minLength(testCases[0].Input)
	}
}

func Benchmark_minLength_string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minLength_string(testCases[0].Input)
	}
}

func Benchmark_minLength_strings_replace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minLength_strings_replace(testCases[0].Input)
	}
}
