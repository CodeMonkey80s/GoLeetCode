package solution3392

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{

	{
		Input:  []int{1, 2, 1, 4, 1},
		Output: 1,
	},
	{
		Input:  []int{1, 1, 1},
		Output: 0,
	},
}

func Test_countSubarrays(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countSubarrays(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countSubarrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countSubarrays(testCases[0].Input)
	}
}
