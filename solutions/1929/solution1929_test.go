package solution1929

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
		Input:  []int{1, 2, 1},
		Output: []int{1, 2, 1, 1, 2, 1},
	},
	{
		Input:  []int{1, 3, 2, 1},
		Output: []int{1, 3, 2, 1, 1, 3, 2, 1},
	},
}

func Test_getConcatenation(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getConcatenation(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_getConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getConcatenation(testCases[0].Input)
	}
}
