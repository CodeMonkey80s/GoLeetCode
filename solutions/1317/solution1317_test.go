package solution1317

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
		Input:  2,
		Output: []int{1, 1},
	},
	{
		Input:  11,
		Output: []int{9, 2},
	},
}

func Test_getNoZeroIntegers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getNoZeroIntegers(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_getNoZeroIntegers(b *testing.B) {
	for b.Loop() {
		_ = getNoZeroIntegers(testCases[0].Input)
	}
}
