package solution451

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "tree",
		Output: "eert",
	},
	{
		Input:  "cccaaa",
		Output: "aaaccc",
	},
	{
		Input:  "Aabb",
		Output: "bbAa",
	},
	{
		Input:  "loveleetcode",
		Output: "eeeelloocdtv",
	},
}

func Test_frequencySort(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := frequencySort(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
