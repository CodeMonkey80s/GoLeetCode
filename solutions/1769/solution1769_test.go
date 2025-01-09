package solution1769

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output []int
}{
	{
		Input:  "110",
		Output: []int{1, 1, 3},
	},
	{
		Input:  "001011",
		Output: []int{11, 8, 5, 4, 3, 4},
	},
}

func Test_minOperations(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minOperations(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
