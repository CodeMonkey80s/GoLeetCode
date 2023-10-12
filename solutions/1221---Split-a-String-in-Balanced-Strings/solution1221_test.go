package solution1221

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "RLRRLLRLRL",
		Output: 4,
	},
	{
		Input:  "RLRRRLLRLL",
		Output: 2,
	},
	{
		Input:  "LLLLRRRR",
		Output: 1,
	},
}

func Test_balancedStringSplit(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := balancedStringSplit(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_balancedStringSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = balancedStringSplit(testCases[0].Input)
	}
}
