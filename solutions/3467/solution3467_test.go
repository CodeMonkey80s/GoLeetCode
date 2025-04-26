package solution3467

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
		Input:  []int{4, 3, 2, 1},
		Output: []int{0, 0, 1, 1},
	},
	{
		Input:  []int{1, 5, 1, 4, 2},
		Output: []int{0, 0, 1, 1, 1},
	},
}

func Test_transformArray(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := transformArray(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
