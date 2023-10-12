package solution1859

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "is2 sentence4 This1 a3",
		Output: "This is a sentence",
	},
	{
		Input:  "Myself2 Me1 I4 and3",
		Output: "Me Myself and I",
	},
}

func Test_sortSentence(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sortSentence(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func Benchmark_sortSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sortSentence(testCases[0].Input)
	}
}
