package solution1310

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB [][]int
	Output []int
}{

	{
		InputA: []int{1, 3, 4, 8},
		InputB: [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}},
		Output: []int{2, 7, 14, 8},
	},
	{
		InputA: []int{4, 8, 2, 10},
		InputB: [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}},
		Output: []int{8, 0, 4, 4},
	},
}

func Test_pivotIndex(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := xorQueries(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
