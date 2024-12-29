package solution2545

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA [][]int
	InputB int
	Output [][]int
}{
	{
		InputA: [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}},
		InputB: 2,
		Output: [][]int{{7, 5, 11, 2}, {10, 6, 9, 1}, {4, 8, 3, 15}},
	},
}

func Test_sortTheStudents(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sortTheStudents(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
