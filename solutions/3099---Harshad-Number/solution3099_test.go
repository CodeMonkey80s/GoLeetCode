package solution3099

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output int
}{
	{
		Input:  18,
		Output: 9,
	},
	{
		Input:  23,
		Output: -1,
	},
}

func Test_sumOfTheDigitsOfHarshadNumber(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sumOfTheDigitsOfHarshadNumber(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_sumOfTheDigitsOfHarshadNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sumOfTheDigitsOfHarshadNumber(testCases[0].Input)
	}
}
