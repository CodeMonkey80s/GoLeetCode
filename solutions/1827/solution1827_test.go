package solution1827

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 1, 1},
		Output: 3,
	},
	{
		Input:  []int{1, 5, 2, 4, 1},
		Output: 14,
	},
	{
		Input:  []int{8},
		Output: 0,
	},
}

func Test_minOperations(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := minOperations(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_minOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minOperations(testCases[1].Input)
	}
}
