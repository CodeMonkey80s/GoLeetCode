package solution890

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB string
	Output []string
}{
	{
		InputA: []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
		InputB: "abb",
		Output: []string{"mee", "aqq"},
	},
	{
		InputA: []string{"a", "b", "c"},
		InputB: "a",
		Output: []string{"a", "b", "c"},
	},
}

func Test_sortVowels(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findAndReplacePattern(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}
