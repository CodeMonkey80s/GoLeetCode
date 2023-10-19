package solution169

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{3, 2, 3},
		Output: 3,
	},
	{
		Input:  []int{2, 2, 1, 1, 1, 2, 2},
		Output: 2,
	},
}

func Test_majorityElement(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := majorityElement(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_majorityElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = majorityElement(testCases[0].Input)
	}
}
