package solution2341

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
		Input:  []int{1, 3, 2, 1, 3, 2, 2},
		Output: []int{3, 1},
	},
	{
		Input:  []int{1, 1},
		Output: []int{1, 0},
	},
	{
		Input:  []int{0},
		Output: []int{0, 1},
	},
}

func Test_numberOfPairs(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numberOfPairs(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_numberOfPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numberOfPairs(testCases[0].Input)
	}
}
