package solution1047

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "abbaca",
		Output: "ca",
	},
	{
		Input:  "azxxzy",
		Output: "ay",
	},
}

func Test_removeDuplicates(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := removeDuplicates(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeDuplicates(testCases[0].Input)
	}
}

func Benchmark_removeDuplicates_string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeDuplicatesStack(testCases[0].Input)
	}
}
