package solution448

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{
	{
		Input:  []int{4, 3, 2, 7, 8, 2, 3, 1},
		Output: []int{5, 6},
	},
	{
		Input:  []int{1, 1},
		Output: []int{2},
	},
}

func Test_findDisappeared(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findDisappearedNumbers(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_find(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findDisappearedNumbers(testCases[0].Input)
	}
}
