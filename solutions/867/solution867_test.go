package solution867

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
		Input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		Output: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
	},
	{
		Input:  [][]int{{1, 2, 3}, {4, 5, 6}},
		Output: [][]int{{1, 4}, {2, 5}, {3, 6}},
	},
}

func Test_transpose(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := transpose(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
