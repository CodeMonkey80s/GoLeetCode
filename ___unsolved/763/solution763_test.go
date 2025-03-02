package solution763

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output []int
}{

	{
		Input:  "ababcbacadefegdehijhklij",
		Output: []int{9, 7, 8},
	},
	{
		Input:  "eccbbbbdec",
		Output: []int{10},
	},
}

func Test_partitionLabels(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := partitionLabels(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
