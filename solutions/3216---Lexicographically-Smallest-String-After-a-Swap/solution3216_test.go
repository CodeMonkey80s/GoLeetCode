package solution3216

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "45320",
		Output: "43520",
	},
	{
		Input:  "001",
		Output: "001",
	},
	{
		Input:  "13",
		Output: "13",
	},
	{
		Input:  "131",
		Output: "113",
	},
}

func Test_countKConstraintSubstrings(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getSmallestString(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
