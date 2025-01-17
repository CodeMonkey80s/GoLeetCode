package solution2903

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	InputC int
	Output []int
}{
	{
		InputA: []int{5, 1, 4, 1},
		InputB: 2,
		InputC: 4,
		Output: []int{0, 3},
	},
	{
		InputA: []int{2, 1},
		InputB: 0,
		InputC: 0,
		Output: []int{0, 0},
	},
	{
		InputA: []int{1, 2, 3},
		InputB: 2,
		InputC: 4,
		Output: []int{-1, -1},
	},
}

func Test_findIndices(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findIndices(tc.InputA, tc.InputB, tc.InputC)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
