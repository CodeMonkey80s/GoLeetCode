package solution3978

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output bool
}{
	{
		Input:  []int{1, 2, 3},
		Output: true,
	},
	{
		Input:  []int{1, 2, 2},
		Output: false,
	},
}

func Test_isMiddleElementUnique(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isMiddleElementUnique(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
