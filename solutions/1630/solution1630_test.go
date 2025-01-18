package solution1630

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	InputC []int
	Output []bool
}{
	{
		InputA: []int{4, 6, 5, 9, 3, 7},
		InputB: []int{0, 0, 2},
		InputC: []int{2, 3, 5},
		Output: []bool{true, false, true},
	},
	{
		InputA: []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10},
		InputB: []int{0, 1, 6, 4, 8, 7},
		InputC: []int{4, 4, 9, 7, 9, 10},
		Output: []bool{false, true, false, false, true, true},
	},
}

func Test_checkArithmeticSubarrays(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := checkArithmeticSubarrays(tc.InputA, tc.InputB, tc.InputC)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_checkArithmeticSubarrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = checkArithmeticSubarrays(testCases[0].InputA, testCases[0].InputB, testCases[0].InputC)
	}
}
