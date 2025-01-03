package solution485

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 1, 0, 1, 1, 1},
		Output: 3,
	},
	{
		Input:  []int{1, 0, 1, 1, 0, 1},
		Output: 2,
	},
}

func Test_findMaxConsecutiveOnes(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findMaxConsecutiveOnes(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findMaxConsecutiveOnes(testCases[0].Input)
	}
}
