package solution26

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
	Total  int
}{
	// Mandatory Test Cases
	{
		Input:  []int{1, 1, 2},
		Output: []int{1, 2},
		Total:  2,
	},
	{
		Input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		Output: []int{0, 1, 2, 3, 4},
		Total:  5,
	},
	// Additional my custom cases
	{
		Input:  []int{1, 1},
		Output: []int{1},
		Total:  1,
	},
}

func Test_removeDuplicates(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: %v, %d\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := removeDuplicates(tc.Input)
			if output != tc.Total {
				t.Errorf("Expected output to be %v but we got %v", tc.Total, output)
			}
		})
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeDuplicates(testCases[0].Input)
	}
}

func Benchmark_removeDuplicates_stdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeDuplicates_stdlib(testCases[0].Input)
	}
}
