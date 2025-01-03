package solution1304

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  int
	Output []int
}{
	{
		Input:  5,
		Output: []int{1, -1, 2, -2, 0},
	},
	{
		Input:  3,
		Output: []int{1, -1, 0},
	},
	{
		Input:  1,
		Output: []int{0},
	},
}

func Test_sumZero(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sumZero(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_sumZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sumZero(testCases[0].Input)
	}
}
