package solution48

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
		Output: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
	},
}

func Test_rotate(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			rotate(tc.Input)
			if !reflect.DeepEqual(tc.Input, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, tc.Input)
			}
		})
	}
}
