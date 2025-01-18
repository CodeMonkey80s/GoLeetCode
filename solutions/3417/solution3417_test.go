package solution3417

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
		Input:  [][]int{{1, 2}, {3, 4}},
		Output: []int{1, 4},
	},
	{
		Input:  [][]int{{2, 1}, {2, 1}, {2, 1}},
		Output: []int{2, 1, 2},
	},
	{
		Input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		Output: []int{1, 3, 5, 7, 9},
	},
}

func Test_zigzagTraversal(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := zigzagTraversal(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
