package solution3498

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{

	{
		Input:  "abc",
		Output: 148,
	},
	{
		Input:  "zaza",
		Output: 160,
	},
}

func Test_reverseDegree(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reverseDegree(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
