package solution2119

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output bool
}{
	{
		Input:  526,
		Output: true,
	},
	{
		Input:  1800,
		Output: false,
	},
	{
		Input:  0,
		Output: true,
	},
}

func Test_isSameAfterReversals(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isSameAfterReversals(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkIsSameAfterReversals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isSameAfterReversals(testCases[0].Input)
	}
}
