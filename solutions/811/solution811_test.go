package solution811

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output []string
}{
	{
		Input:  []string{"9001 discuss.leetcode.com"},
		Output: []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
	},
}

func Test_subdomainVisits(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			m1 := make(map[int]int)
			m2 := make(map[int]int)
			output := subdomainVisits(tc.Input)
			seq1 := slices.All(output)
			maps.Collect(seq1)
			seq2 := slices.All(tc.Output)
			maps.Collect(seq2)
			if !reflect.DeepEqual(m1, m2) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
