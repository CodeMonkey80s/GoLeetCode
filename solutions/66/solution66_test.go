package solution66

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{
	// Mandatory Test Cases
	{
		Input:  []int{1, 2, 3},
		Output: []int{1, 2, 4},
	},
	{
		Input:  []int{4, 3, 2, 1},
		Output: []int{4, 3, 2, 2},
	},
	{
		Input:  []int{9},
		Output: []int{1, 0},
	},
	// Additional my custom cases
	{
		Input:  []int{8, 9, 9, 9},
		Output: []int{9, 0, 0, 0},
	},
}

func Test_plusOne(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := plusOne(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_plusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = plusOne(testCases[0].Input)
	}
}
