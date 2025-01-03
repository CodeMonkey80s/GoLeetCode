package solution1282

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output [][]int
}{
	{
		Input:  []int{3, 3, 3, 3, 3, 1, 3},
		Output: [][]int{{5}, {0, 1, 2}, {3, 4, 6}},
	},
	{
		Input:  []int{2, 1, 3, 3, 3, 2},
		Output: [][]int{{1}, {0, 5}, {2, 3, 4}},
	},
}

func Test_groupThePeople(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := groupThePeople(tc.Input)
			slices.SortFunc(output, func(a []int, b []int) int {
				return len(a) - len(b)
			})
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_groupThePeople(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = groupThePeople(testCases[0].Input)
	}
}
