package solution830

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output [][]int
}{
	// Mandatory Test Cases
	{
		Input:  "abbxxxxzzy",
		Output: [][]int{{3, 6}},
	},
	{
		Input:  "abc",
		Output: [][]int{},
	},
	{
		Input:  "abcdddeeeeaabbbcd",
		Output: [][]int{{3, 5}, {6, 9}, {12, 14}},
	},
	// Additional my custom cases
	{
		Input:  "aaa",
		Output: [][]int{{0, 2}},
	},
	{
		Input:  "babaaaabbb",
		Output: [][]int{{3, 6}, {7, 9}},
	},
	{
		Input:  "nnnhaaannnm",
		Output: [][]int{{0, 2}, {4, 6}, {7, 9}},
	},
}

func Test_largeGroupPositions(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := largeGroupPositions(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_largeGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = largeGroupPositions(testCases[0].Input)
	}
}
