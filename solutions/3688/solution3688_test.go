package solution3688

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
		Input:  []int{1, 2, 3, 4, 5, 6},
		Output: 6,
	},
	{
		Input:  []int{7, 9, 11},
		Output: 0,
	},
	{
		Input:  []int{1, 8, 16},
		Output: 24,
	},
}

func Test_evenNumbersBitwiseORs(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := evenNumberBitwiseORs(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
