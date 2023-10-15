package solution1313

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
		Input:  []int{1, 2, 3, 4},
		Output: []int{2, 4, 4, 4},
	},
	{
		Input:  []int{1, 1, 2, 3},
		Output: []int{1, 3, 3},
	},
}

func Test_decompress(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := decompressRLElist(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_decompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = decompressRLElist(testCases[0].Input)
	}
}
