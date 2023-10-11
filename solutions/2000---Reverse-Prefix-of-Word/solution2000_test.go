package solution2000

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB byte
	Output string
}{
	{
		InputA: "abcdefd",
		InputB: 'd',
		Output: "dcbaefd",
	},
	{
		InputA: "xyxzxe",
		InputB: 'z',
		Output: "zxyxxe",
	},
	{
		InputA: "abcd",
		InputB: 'z',
		Output: "abcd",
	},
}

func Test_reversePrefix(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reversePrefix(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_reversePrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reversePrefix(testCases[0].InputA, testCases[0].InputB)
	}
}
