package solution2914

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
		Input:  "1001",
		Output: 2,
	},
	{
		Input:  "10",
		Output: 1,
	},
	{
		Input:  "0000",
		Output: 0,
	},
}

func Test_minChanges(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minChanges(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
