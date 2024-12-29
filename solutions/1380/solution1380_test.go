package solution1380

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output []int
}{
	{
		Input: [][]int{
			{3, 7, 8},
			{9, 11, 13},
			{15, 16, 17},
		},
		Output: []int{15},
	},
	{
		Input: [][]int{
			{1, 10, 4, 2},
			{9, 3, 8, 7},
			{15, 16, 17, 12},
		},
		Output: []int{12},
	},
	{
		Input: [][]int{
			{7, 8},
			{1, 2},
		},
		Output: []int{7},
	},
}

func Test_luckyNumbers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := luckyNumbers(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_luckyNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = luckyNumbers(testCases[1].Input)
	}
}
