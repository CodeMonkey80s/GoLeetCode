package solution2221

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 2, 3, 4, 5},
		Output: 8,
	},
	{
		Input:  []int{5},
		Output: 5,
	},
}

func Test_triangularSum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := triangularSum(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_triangularSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = triangularSum(testCases[0].Input)
	}
}
