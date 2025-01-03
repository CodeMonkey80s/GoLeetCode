package solution2595

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
		Input:  17,
		Output: []int{2, 0},
	},
	{
		Input:  2,
		Output: []int{0, 1},
	},
	{
		Input:  4,
		Output: []int{1, 0},
	},
}

func Test_evenOddBit(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := evenOddBit(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_evenOddBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = evenOddBit(testCases[0].Input)
	}
}

func Benchmark_evenOddBit_first_approach(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = evenOddBit_first_approach(testCases[0].Input)
	}
}
