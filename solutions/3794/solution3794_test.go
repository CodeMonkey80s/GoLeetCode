package solution3794

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB int
	Output string
}{
	{
		InputA: "abcd",
		InputB: 2,
		Output: "bacd",
	},
	{
		InputA: "xyz",
		InputB: 3,
		Output: "zyx",
	},
	{
		InputA: "hey",
		InputB: 1,
		Output: "hey",
	},
}

func Test_reversePrefix(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reversePrefix(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
