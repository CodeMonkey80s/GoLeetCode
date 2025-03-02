package solution3314

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
		Input:  []int{2, 3, 5, 7},
		Output: []int{-1, 1, 4, 3},
	},
	{
		Input:  []int{11, 13, 31},
		Output: []int{9, 12, 15},
	},
}

func Test_minBitwiseArray(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minBitwiseArray(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
