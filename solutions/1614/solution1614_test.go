package solution1614

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "(1+(2*3)+((8)/4))+1",
		Output: 3,
	},
	{
		Input:  "(1)+((2))+(((3)))",
		Output: 3,
	},
}

func Test_maxDepth(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maxDepth(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maxDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxDepth(testCases[0].Input)
	}
}
