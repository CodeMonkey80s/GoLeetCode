package solution1122

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	Output []int
}{
	{
		InputA: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
		InputB: []int{2, 1, 4, 3, 9, 6},
		Output: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
	},
	{
		InputA: []int{28, 6, 22, 8, 44, 17},
		InputB: []int{22, 28, 8, 6},
		Output: []int{22, 28, 8, 6, 17, 44},
	},
}

func Test_relativeSortArray(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := relativeSortArray(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
