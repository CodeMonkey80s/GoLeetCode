package solution2932

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 2, 3, 4, 5},
		Output: 7,
	},
	{
		Input:  []int{10, 100},
		Output: 0,
	},
	{
		Input:  []int{5, 6, 25, 30},
		Output: 7,
	},
}

func Test_maximumStrongPairXor(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maximumStrongPairXor(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkMaximumStrongPairXor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maximumStrongPairXor(testCases[0].Input)
	}
}
