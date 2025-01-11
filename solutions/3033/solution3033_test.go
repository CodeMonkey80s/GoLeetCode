package solution3033

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output [][]int
}{
	{
		Input:  [][]int{{1, 2, -1}, {4, -1, 6}, {7, 8, 9}},
		Output: [][]int{{1, 2, 9}, {4, 8, 6}, {7, 8, 9}},
	},
	{
		Input:  [][]int{{3, -1}, {5, 2}},
		Output: [][]int{{3, 2}, {5, 2}},
	},
	{
		Input:  [][]int{{-1, 0, 0, 2, 2}, {2, 0, 0, 2, 1}, {4, 3, 2, 1, 1}, {-1, -1, 0, 2, 4}, {1, 0, 3, -1, 0}},
		Output: [][]int{{4, 0, 0, 2, 2}, {2, 0, 0, 2, 1}, {4, 3, 2, 1, 1}, {4, 3, 0, 2, 4}, {1, 0, 3, 2, 0}},
	},
}

func Test_modifiedMatrix(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := modifiedMatrix(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
