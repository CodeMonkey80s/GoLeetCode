package solution916

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB []string
	Output []string
}{

	// {
	// 	InputA: []string{"amazon", "apple", "facebook", "google", "leetcode"},
	// 	InputB: []string{"e", "o"},
	// 	Output: []string{"facebook", "google", "leetcode"},
	// },
	// {
	// 	InputA: []string{"amazon", "apple", "facebook", "google", "leetcode"},
	// 	InputB: []string{"l", "e"},
	// 	Output: []string{"apple", "google", "leetcode"},
	// },
	{
		InputA: []string{"amazon", "apple", "facebook", "google", "leetcode"},
		InputB: []string{"lo", "eo"},
		Output: []string{"google", "leetcode"},
	},
}

func Test_wordSubsets(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := wordSubsets(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
