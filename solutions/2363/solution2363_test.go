package solution2363

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA [][]int
	InputB [][]int
	Output [][]int
}{
	{
		InputA: [][]int{{1, 1}, {4, 5}, {3, 8}},
		InputB: [][]int{{3, 1}, {1, 5}},
		Output: [][]int{{1, 6}, {3, 9}, {4, 5}},
	},
	{
		InputA: [][]int{{1, 1}, {3, 2}, {2, 3}},
		InputB: [][]int{{2, 1}, {3, 2}, {1, 3}},
		Output: [][]int{{1, 4}, {2, 4}, {3, 4}},
	},
	{
		InputA: [][]int{{1, 3}, {2, 2}},
		InputB: [][]int{{7, 1}, {2, 2}, {1, 4}},
		Output: [][]int{{1, 7}, {2, 4}, {7, 1}},
	},
}

func Test_numberOfPairs(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := mergeSimilarItems(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
