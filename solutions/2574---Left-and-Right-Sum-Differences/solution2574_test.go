package solution2574

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
		Input:  []int{10, 4, 8, 3},
		Output: []int{15, 1, 11, 22},
	},
	{
		Input:  []int{1},
		Output: []int{0},
	},
}

func Test_leftRightDifference(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := leftRightDifference(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_leftRightDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = leftRightDifference(testCases[0].Input)
	}
}
