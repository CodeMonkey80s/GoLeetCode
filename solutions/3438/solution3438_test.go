package solution3438

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "2523533",
		Output: "23",
	},
	{
		Input:  "221",
		Output: "21",
	},
	{
		Input:  "22",
		Output: "",
	},
}

func Test_findValidPair(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findValidPair(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_makeGood(b *testing.B) {
	for b.Loop() {
		_ = findValidPair(testCases[0].Input)
	}
}
