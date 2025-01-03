package solution70

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  2,
		Output: 2,
	},
	{
		Input:  3,
		Output: 3,
	},
	// Additional my custom cases
}

func Test_climbStairs(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := climbStairs(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_climbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = climbStairs(testCases[0].Input)
	}
}
