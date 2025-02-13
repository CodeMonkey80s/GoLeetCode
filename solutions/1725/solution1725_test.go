package solution1725

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
		Input:  [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}},
		Output: 3,
	},
	{
		Input:  [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}},
		Output: 3,
	},
}

func Test_countGoodRectangles(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countGoodRectangles(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countGoodRectangles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countGoodRectangles(testCases[0].Input)
	}
}
