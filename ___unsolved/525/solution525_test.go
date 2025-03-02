package solution525

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{0, 1},
		Output: 2,
	},
	{
		Input:  []int{0, 1, 0},
		Output: 2,
	},
	{
		Input:  []int{1, 0, 0, 0, 1, 1, 1, 0},
		Output: 6,
	},
	{
		Input:  []int{0, 1, 0, 1},
		Output: 4,
	},
	{
		Input:  []int{0, 1, 1},
		Output: 2,
	},
	{
		Input:  []int{1, 1, 0, 0, 1, 1, 0, 1, 1},
		Output: 0,
	},
}

func Test_findMaxLength(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findMaxLength(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
