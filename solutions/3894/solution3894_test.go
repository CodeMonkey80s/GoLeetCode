package solution3894

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  int
	Output string
}{
	{
		Input:  60,
		Output: "Red",
	},
	{
		Input:  30,
		Output: "Orange",
	},
	{
		Input:  0,
		Output: "Green",
	},
	{
		Input:  5,
		Output: "Invalid",
	},
}

func Test_trafficSignal(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := trafficSignal(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
