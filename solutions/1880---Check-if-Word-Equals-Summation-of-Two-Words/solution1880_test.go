package solution1880

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	InputC string
	Output bool
}{
	{
		InputA: "acb",
		InputB: "cba",
		InputC: "cdb",
		Output: true,
	},
	{
		InputA: "aaa",
		InputB: "a",
		InputC: "aab",
		Output: false,
	},
	{
		InputA: "aaa",
		InputB: "a",
		InputC: "aaaa",
		Output: true,
	},
}

func Test_isSumEqual(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v %v Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isSumEqual(tc.InputA, tc.InputB, tc.InputC)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isSumEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isSumEqual(testCases[0].InputA, testCases[0].InputB, testCases[0].InputC)
	}
}
