package solution1021

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{

	{
		Input:  "(()())(())",
		Output: "()()()",
	},
	{
		Input:  "(()())(())(()(()))",
		Output: "()()()()(())",
	},
	{

		Input:  "()()",
		Output: "",
	},
}

func Test_removeOuterParentheses(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := removeOuterParentheses(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_removeOuterParentheses(b *testing.B) {
	for b.Loop() {
		_ = removeOuterParentheses(testCases[0].Input)
	}
}
