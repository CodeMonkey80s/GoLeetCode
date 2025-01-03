package solution746

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  []int{10, 15, 20},
		Output: 15,
	},
	{
		Input:  []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		Output: 6,
	},
	// Additional my custom cases
}

func Test_cost(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minCostClimbingStairs(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_cost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minCostClimbingStairs(testCases[0].Input)
	}
}
