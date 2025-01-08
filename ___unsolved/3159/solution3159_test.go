package solution3159

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	InputC int
	Output []int
}{
	{
		InputA: []int{1, 3, 1, 7},
		InputB: []int{1, 3, 2, 4},
		InputC: 1,
		Output: []int{0, -1, 2, -1},
	},
	{
		InputA: []int{1, 2, 3},
		InputB: []int{10},
		InputC: 5,
		Output: []int{-1},
	},
}

func Test_occurrencesOfElement(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := occurrencesOfElement(tc.InputA, tc.InputB, tc.InputC)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
