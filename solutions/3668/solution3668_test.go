package solution3668

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
		InputA: []int{3, 1, 2, 5, 4},
		InputB: []int{1, 3, 4},
		Output: []int{3, 1, 4},
	},
	{
		InputA: []int{1, 4, 5, 3, 2},
		InputB: []int{2, 5},
		Output: []int{5, 2},
	},
}

func Test_recoverOrder(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := recoverOrder(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(tc.Output, output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_gcdOfOddEvenSums(b *testing.B) {
	for b.Loop() {
		_ = recoverOrder(testCases[0].InputA, testCases[0].InputB)
	}
}
