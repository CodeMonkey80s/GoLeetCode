package solution1309

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "10#11#12",
		Output: "jkab",
	},
	{
		Input:  "1326#",
		Output: "acz",
	},
}

func Test_freqAlphabets(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := freqAlphabets(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_freqAlphabet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = freqAlphabets(testCases[0].Input)
	}
}
