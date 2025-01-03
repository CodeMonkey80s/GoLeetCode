package solution2778

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
		Input:  []int{1, 2, 3, 4},
		Output: 21,
	},
	{
		Input:  []int{2, 7, 1, 19, 18, 3},
		Output: 63,
	},
}

func Test_sumOfSquares(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sumOfSquares(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_sumOfSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sumOfSquares(testCases[0].Input)
	}
}
