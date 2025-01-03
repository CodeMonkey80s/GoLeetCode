package solution1480

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
		Input:  []int{1, 2, 3, 4},
		Output: []int{1, 3, 6, 10},
	},
	{
		Input:  []int{1, 1, 1, 1, 1},
		Output: []int{1, 2, 3, 4, 5},
	},
	{
		Input:  []int{3, 1, 2, 10, 1},
		Output: []int{3, 4, 6, 16, 17},
	},
}

func Test_runningSum(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := runningSum(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_runningSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = runningSum(testCases[0].Input)
	}
}
