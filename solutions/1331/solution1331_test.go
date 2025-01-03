package solution1331

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{
	{
		Input:  []int{40, 10, 20, 30},
		Output: []int{4, 1, 2, 3},
	},
	{
		Input:  []int{100, 100, 100},
		Output: []int{1, 1, 1},
	},
	{
		Input:  []int{37, 12, 28, 9, 100, 56, 80, 5, 12},
		Output: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
	},
}

func Test_arrayRankTransform(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := arrayRankTransform(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_arrayRankTransform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = arrayRankTransform(testCases[0].Input)
	}
}
