package solution46

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output [][]int
}{
	{
		Input:  []int{1, 2, 3},
		Output: [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}},
	},
	{
		Input:  []int{0, 1},
		Output: [][]int{{0, 1}, {1, 0}},
	},
	{
		Input:  []int{1},
		Output: [][]int{{1}},
	},
}

func Test_permute(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := permute(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
