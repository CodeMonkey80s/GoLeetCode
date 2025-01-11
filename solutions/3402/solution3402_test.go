package solution3402

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output int
}{
	{
		Input:  [][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}},
		Output: 15,
	},
	{
		Input:  [][]int{{3, 2, 1}, {2, 1, 0}, {1, 2, 3}},
		Output: 12,
	},
}

func Test_minimumOperations(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minimumOperations(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minimumOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minimumOperations(testCases[0].Input)
	}
}
