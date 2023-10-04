package solution944

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  []string{"cba", "daf", "ghi"},
		Output: 1,
	},
	{
		Input:  []string{"a", "b"},
		Output: 0,
	},
	{
		Input:  []string{"zyx", "wvu", "tsr"},
		Output: 3,
	},
	// Additional my custom cases
	{
		Input:  []string{"rrjk", "furt", "guzm"},
		Output: 2,
	},
}

func Test_minDeletionSize(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minDeletionSize(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %d but we got %d", tc.Output, output)
			}
		})
	}
}

func Benchmark_minDeletionSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minDeletionSize(testCases[0].Input)
	}
}
