package solution3908

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA int
	InputB int
	Output bool
}{
	{
		InputA: 101,
		InputB: 0,
		Output: true,
	},
	{
		InputA: 232,
		InputB: 2,
		Output: false,
	},
	{
		InputA: 5,
		InputB: 1,
		Output: false,
	},
}

func Test_validDigit(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := validDigit(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
