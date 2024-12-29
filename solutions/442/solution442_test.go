package solution442

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
		Output: []int{2, 3},
	},
	{
		Input:  []int{1, 1, 2},
		Output: []int{1},
	},
	{
		Input:  []int{1},
		Output: []int{},
	},
}

func Test_findDuplicates(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findDuplicates(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findDuplicates(testCases[0].Input)
	}
}
