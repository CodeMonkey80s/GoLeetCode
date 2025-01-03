package solution121

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
		Input:  []int{7, 1, 5, 3, 6, 4},
		Output: 5,
	},
	{
		Input:  []int{7, 6, 4, 3, 1},
		Output: 0,
	},
	// Additional my custom cases
	{
		Input:  []int{2, 4, 1},
		Output: 2,
	},
	{
		Input:  []int{1},
		Output: 0,
	},
	{
		Input:  []int{3, 2, 6, 5, 0, 3},
		Output: 4,
	},
}

func Test_maxProfit(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maxProfit(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_removeElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxProfit(testCases[0].Input)
	}
}
