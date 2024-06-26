package solution3120

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "aaAbcBC",
		Output: 3,
	},
	{
		Input:  "abc",
		Output: 0,
	},
	{
		Input:  "abBCab",
		Output: 1,
	},
	{
		Input:  "Z",
		Output: 0,
	},
	{
		Input:  "z",
		Output: 0,
	},
	{
		Input:  "BBbab",
		Output: 1,
	},
	{
		Input:  "i",
		Output: 0,
	},
}

func Test_numberOfSpecialCharacters(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numberOfSpecialCharacters(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isArraySpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numberOfSpecialCharacters(testCases[2].Input)
	}
}
