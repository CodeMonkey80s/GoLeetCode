package solution1399

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  int
	Output int
}{
	{
		Input:  13,
		Output: 4,
	},
	{
		Input:  2,
		Output: 2,
	},
	{
		Input:  24,
		Output: 5,
	},
	{
		Input:  33,
		Output: 4,
	},
}

func Test_countLargestGroup(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countLargestGroup(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countLargestGroup(b *testing.B) {
	for b.Loop() {
		_ = countLargestGroup(testCases[0].Input)
	}
}
