package solution3005

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 2, 2, 3, 1, 4},
		Output: 4,
	},
	{
		Input:  []int{1, 2, 3, 4, 5},
		Output: 5,
	},
}

func Test_maxFrequencyElements(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maxFrequencyElements(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maxFrequencyElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxFrequencyElements(testCases[0].Input)
	}
}
