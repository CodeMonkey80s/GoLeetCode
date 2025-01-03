package solution3310

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "hello",
		Output: 13,
	},
	{
		Input:  "zaz",
		Output: 50,
	},
}

func Test_scoreOfScreens(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := scoreOfString(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isArraySpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = scoreOfString(testCases[0].Input)
	}
}
