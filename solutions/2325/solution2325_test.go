package solution2325

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output string
}{
	{
		InputA: "the quick brown fox jumps over the lazy dog",
		InputB: "vkbs bs t suepuv",
		Output: "this is a secret",
	},
	{
		InputA: "eljuxhpwnyrdgtqkviszcfmabo",
		InputB: "zwx hnfx lqantp mnoeius ycgk vcnjrdb",
		Output: "the five boxing wizards jump quickly",
	},
}

func Test_decodeMessage(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := decodeMessage(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_decodeMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = decodeMessage(testCases[0].InputA, testCases[0].InputB)
	}
}
