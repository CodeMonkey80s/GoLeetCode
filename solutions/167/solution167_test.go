package solution167

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output []int
}{
	{
		InputA: []int{2, 7, 11, 15},
		InputB: 9,
		Output: []int{1, 2},
	},
	{
		InputA: []int{2, 3, 4},
		InputB: 6,
		Output: []int{1, 3},
	},
	{
		InputA: []int{-1, 0},
		InputB: -1,
		Output: []int{1, 2},
	},
}

func Test_twoSum(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := twoSum(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = twoSum(testCases[0].InputA, testCases[0].InputB)
	}
}
