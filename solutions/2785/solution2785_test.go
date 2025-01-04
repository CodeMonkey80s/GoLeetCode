package solution2785

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "lEetcOde",
		Output: "lEOtcede",
	},
	{
		Input:  "lYmpH",
		Output: "lYmpH",
	},
}

func Test_sortVowels(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sortVowelsV2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func Benchmark_sortVowelsV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sortVowelsV2(testCases[0].Input)
	}
}

func Benchmark_sortVowelsV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sortVowelsV1(testCases[0].Input)
	}
}
