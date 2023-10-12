package solution1512

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 2, 3, 1, 1, 3},
		Output: 4,
	},
	{
		Input:  []int{1, 1, 1, 1},
		Output: 6,
	},
	{
		Input:  []int{1, 2, 3},
		Output: 0,
	},
}

func Test_numIdenticalPairs(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numIdenticalPairs(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_squareIsWhite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numIdenticalPairs(testCases[0].Input)
	}
}
