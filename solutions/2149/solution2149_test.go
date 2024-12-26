package solution2149

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
		Input:  []int{3, 1, -2, -5, 2, -4},
		Output: []int{3, -2, 1, -5, 2, -4},
	},
}

func Test_rearrangeArray(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := rearrangeArray(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_rearrangeArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rearrangeArray(testCases[0].Input)
	}
}
