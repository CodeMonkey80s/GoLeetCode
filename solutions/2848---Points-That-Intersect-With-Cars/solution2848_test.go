package solution2848

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
		Input: [][]int{
			{3, 6},
			{1, 5},
			{4, 7},
		},
		Output: 7,
	},
	{

		Input: [][]int{
			{1, 3},
			{5, 8},
		},
		Output: 7,
	},
}

func Test_numberOfPoints(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numberOfPoints(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_numberOfPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numberOfPoints(testCases[0].Input)
	}
}
