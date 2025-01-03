package solution1365

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
		Input:  []int{8, 1, 2, 2, 3},
		Output: []int{4, 0, 1, 1, 3},
	},
	{
		Input:  []int{6, 5, 4, 8},
		Output: []int{2, 1, 0, 3},
	},
	{
		Input:  []int{7, 7, 7, 7},
		Output: []int{0, 0, 0, 0},
	},
}

func Test_smallerNumbers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := smallerNumbersThanCurrent(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_smallerNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = smallerNumbersThanCurrent(testCases[0].Input)
	}
}
