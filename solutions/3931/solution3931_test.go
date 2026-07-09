package solution3931

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	{
		Input:  "132",
		Output: true,
	},
	{
		Input:  "129",
		Output: false,
	},
}

func Test_isAdjacentDiffAtMostTwo(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isAdjacentDiffAtMostTwo(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
