package solution303

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB [][]int
	Output []string
}{
	{
		InputA: []string{"NumArray", "sumRange", "sumRange", "sumRange"},
		InputB: [][]int{{-2, 0, 3, -5, 2, -1}, {0, 2}, {2, 5}, {0, 5}},
		Output: []string{"null", "1", "-1", "-3"},
	},
}

func Test_sumRange(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			var output []string
			obj := Constructor(tc.InputB[0])
			output = append(output, "null")
			for i, cmd := range tc.InputA {
				switch cmd {
				case "sumRange":
					v := obj.SumRange(tc.InputB[i][0], tc.InputB[i][1])
					output = append(output, strconv.Itoa(v))
				}
			}
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
