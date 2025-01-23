package solution645

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
		Input:  []int{1, 2, 2, 4},
		Output: []int{2, 3},
	},
	{
		Input:  []int{1, 1},
		Output: []int{1, 2},
	},
}

func Test_findErrorNums(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findErrorNums(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findErrorNums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findErrorNums(testCases[0].Input)
	}
}
