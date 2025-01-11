package solution3379

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
		Input:  []int{3, -2, 1, 1},
		Output: []int{1, 1, 1, 3},
	},
	{
		Input:  []int{-1, 4, -1},
		Output: []int{-1, -1, 4},
	},
	{
		Input:  []int{-10, -10},
		Output: []int{-10, -10},
	},
}

func Test_constructTransformedArray(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := constructTransformedArray(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_constructTransformedArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = constructTransformedArray(testCases[0].Input)
	}
}
