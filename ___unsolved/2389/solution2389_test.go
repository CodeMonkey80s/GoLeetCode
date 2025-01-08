package solution2389

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	Output []int
}{
	{
		InputA: []int{4, 5, 2, 1},
		InputB: []int{3, 10, 21},
		Output: []int{2, 3, 4},
	},
	{
		InputA: []int{2, 3, 4, 5},
		InputB: []int{1},
		Output: []int{0},
	},
}

func Test_answerQueries(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := answerQueries(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
