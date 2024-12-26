package solution2482

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
		Input:  [][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}},
		Output: [][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}},
	},
	{
		Input:  [][]int{{1, 1, 1}, {1, 1, 1}},
		Output: [][]int{{5, 5, 5}, {5, 5, 5}},
	},
}

func Test_findMatrix(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := onesMinusZeros(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
