package solution3536

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
		Input:  31,
		Output: 3,
	},
	{
		Input:  22,
		Output: 4,
	},
	{
		Input:  124,
		Output: 8,
	},
}

func Test_maxProduct(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maxProduct(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_makeGood(b *testing.B) {
	for b.Loop() {
		_ = maxProduct(testCases[0].Input)
	}
}
