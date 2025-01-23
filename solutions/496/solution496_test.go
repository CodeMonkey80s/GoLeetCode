package solution496

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	Output []int
}{
	{
		InputA: []int{4, 1, 2},
		InputB: []int{1, 3, 4, 2},
		Output: []int{-1, 3, -1},
	},
	{
		InputA: []int{2, 4},
		InputB: []int{1, 2, 3, 4},
		Output: []int{3, -1},
	},
	{
		InputA: []int{1, 3, 5, 2, 4},
		InputB: []int{6, 5, 4, 3, 2, 1, 7},
		Output: []int{7, 7, 7, 7, 7},
	},
}

func Test_nextGreaterElement(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := nextGreaterElementV2(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_nextGreaterElementV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = nextGreaterElementV2(testCases[2].InputA, testCases[2].InputB)
	}
}
