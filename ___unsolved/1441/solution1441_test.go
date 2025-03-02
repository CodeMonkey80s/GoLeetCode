package solution1441

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output []string
}{
	{
		InputA: []int{1, 3},
		InputB: 3,
		Output: []string{"Push", "Push", "Pop", "Push"},
	},
	{
		InputA: []int{1, 2, 3},
		InputB: 3,
		Output: []string{"Push", "Push", "Push"},
	},
	{
		InputA: []int{1, 2},
		InputB: 4,
		Output: []string{"Push", "Push"},
	},
	{
		InputA: []int{2, 3, 4},
		InputB: 4,
		Output: []string{"Push", "Pop", "Push", "Push", "Push"},
	},
}

func Test_countDistinctIntegers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := buildArray(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
