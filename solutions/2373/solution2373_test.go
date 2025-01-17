package solution2373

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
		Input:  [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}},
		Output: [][]int{{9, 9}, {8, 6}},
	},
	{
		Input:  [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
		Output: [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
	},
}

func Test_largestLocal(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := largestLocal(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
