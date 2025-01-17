package solution2037

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	Output int
}{
	{
		InputA: []int{3, 1, 5},
		InputB: []int{2, 7, 4},
		Output: 4,
	},
	{
		InputA: []int{4, 1, 5, 9},
		InputB: []int{1, 3, 2, 6},
		Output: 7,
	},
	{
		InputA: []int{2, 2, 6, 6},
		InputB: []int{1, 3, 2, 6},
		Output: 4,
	},
}

func Test_minMovesToSeat(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minMovesToSeat(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
