package solution1897

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output bool
}{
	{
		Input:  []string{"abc", "aabc", "bc"},
		Output: true,
	},
	{
		Input:  []string{"ab", "a"},
		Output: false,
	},

	{
		Input:  []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaa", "aaaaaaaa", "a", "a", "a", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "bbbbbbbbbbbb", "bbbbbbbbbbb", "bbbbbbbbbbbbbbbbb", "bbbbbbbbbbbbbbbbbbb", "bbb", "bb", "b", "ccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc", "ccccccccc", "cccccccccccccccc", "ccccccc", "ccccccccc", "ccccc", "ccccccccccccccccccc", "cccccccc", "dddddddddddddddddd", "dddddddddd", "dddddddddddddd", "ddddd", "d", "d", "d", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "eeeeeeeeeeeeeeeeee", "eeeeeeeeeeeeeeeeeeeeeeeeeee", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "eee", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "e", "e", "eeee", "eeee", "ee", "eeee", "e", "ee"},
		Output: true,
	},
}

func Test_makeEqual(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := makeEqual(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
