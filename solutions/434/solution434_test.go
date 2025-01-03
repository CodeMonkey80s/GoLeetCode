package solution434

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "Hello, my name is John",
		Output: 5,
	},
	{
		Input:  "Hello",
		Output: 1,
	},
	{
		Input:  "",
		Output: 0,
	},
	{
		Input:  "           ",
		Output: 0,
	},
}

func Test_countSegments(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countSegments(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_missingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countSegments(testCases[0].Input)
	}
}

func Benchmark_missingNumber_strings_fields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countSegments_strings_fields(testCases[0].Input)
	}
}
