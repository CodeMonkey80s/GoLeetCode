package solution75

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{

	{
		Input:  []int{2, 0, 2, 1, 1, 0},
		Output: []int{0, 0, 1, 1, 2, 2},
	},
	{
		Input:  []int{2, 0, 1},
		Output: []int{0, 1, 2},
	},
}

func Test_sortColors(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			sortColors(tc.Input)
			if !reflect.DeepEqual(tc.Input, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, tc.Input)
			}
		})
	}
}
