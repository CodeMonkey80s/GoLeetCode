package solution3386

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output int
}{
	{
		Input:  [][]int{{1, 2}, {2, 5}, {3, 9}, {1, 15}},
		Output: 1,
	},
	{
		Input:  [][]int{{10, 5}, {1, 7}},
		Output: 10,
	},
	{
		Input:  [][]int{{10, 4}, {1, 6}, {7, 14}},
		Output: 7,
	},
	{
		Input:  [][]int{{5, 4}, {20, 14}},
		Output: 20,
	},
	{
		Input:  [][]int{{1, 4}, {18, 5}, {15, 7}, {12, 9}, {1, 11}, {18, 13}, {16, 17}},
		Output: 1,
	},
	{
		Input:  [][]int{{7, 1}, {19, 3}, {9, 4}, {12, 5}, {2, 8}, {15, 10}, {18, 12}, {7, 14}, {19, 16}},
		Output: 2,
	},
}

func Test_buttonWithLongestTime(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := buttonWithLongestTime(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_buttonWithLongestTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = buttonWithLongestTime(testCases[0].Input)
	}
}
