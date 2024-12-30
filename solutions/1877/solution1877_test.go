package solution1877

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
		Input:  []int{3, 5, 2, 3},
		Output: 7,
	},
	{
		Input:  []int{3, 5, 4, 2, 4, 6},
		Output: 8,
	},
	{
		Input:  []int{3, 2, 4, 1, 1, 5, 1, 3, 5, 1},
		Output: 6,
	},

	{
		Input:  []int{4, 1, 5, 1, 2, 5, 1, 5, 5, 4},
		Output: 8,
	},
}

func Test_minPairSum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minPairSum(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minPairSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minPairSum(testCases[2].Input)
	}
}
