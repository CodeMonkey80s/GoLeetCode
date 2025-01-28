package solution1652

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output []int
}{
	{
		InputA: []int{5, 7, 1, 4},
		InputB: 3,
		Output: []int{12, 10, 16, 13},
	},
	{
		InputA: []int{1, 2, 3, 4},
		InputB: 0,
		Output: []int{0, 0, 0, 0},
	},
	{
		InputA: []int{2, 4, 9, 3},
		InputB: -2,
		Output: []int{12, 5, 6, 13},
	},
}

func Test_decrypt(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := decrypt(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
