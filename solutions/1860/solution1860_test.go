package solution1860

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA int
	InputB int
	Output []int
}{
	{
		InputA: 2,
		InputB: 2,
		Output: []int{3, 1, 0},
	},
}

func Test_memLeak(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := memLeak(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
