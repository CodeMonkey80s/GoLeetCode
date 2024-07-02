package solution3162

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	InputC int
	Output int
}{
	{
		InputA: []int{1, 3, 4},
		InputB: []int{1, 3, 4},
		InputC: 1,
		Output: 5,
	},
	{
		InputA: []int{1, 2, 4, 12},
		InputB: []int{2, 4},
		InputC: 3,
		Output: 2,
	},
}

func Test_numberOfPairs(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numberOfPairs(tc.InputA, tc.InputB, tc.InputC)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
