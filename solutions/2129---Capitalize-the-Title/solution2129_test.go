package solution2129

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "capiTalIze tHe titLe",
		Output: "Capitalize The Title",
	},
	{
		Input:  "First leTTeR of EACH Word",
		Output: "First Letter of Each Word",
	},
	{
		Input:  "i lOve leetcode",
		Output: "i Love Leetcode",
	},
	{
		Input:  "L hV",
		Output: "l hv",
	},
}

func Test_capitalizeTitle(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := capitalizeTitle(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func Benchmark_capitalizeTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = capitalizeTitle(testCases[0].Input)
	}
}
