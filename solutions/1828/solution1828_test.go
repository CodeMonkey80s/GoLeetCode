package solution1828

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA [][]int
	InputB [][]int
	Output []int
}{
	{
		InputA: [][]int{
			{1, 3}, {3, 3}, {5, 3}, {2, 2},
		},
		InputB: [][]int{
			{2, 3, 1}, {4, 3, 1}, {1, 1, 2},
		},
		Output: []int{3, 2, 2},
	},
	{
		InputA: [][]int{
			{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5},
		},
		InputB: [][]int{
			{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3},
		},
		Output: []int{2, 3, 2, 4},
	},
}

func Test_countPoints(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countPoints(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countPoints(testCases[0].InputA, testCases[0].InputB)
	}
}
