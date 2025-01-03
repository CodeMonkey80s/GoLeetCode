package solution268

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{3, 0, 1},
		Output: 2,
	},
}

func Test_missingNumber(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := missingNumber(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_missingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		missingNumber(testCases[0].Input)
	}
}
