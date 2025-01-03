package solution1431

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output []bool
}{
	{
		InputA: []int{2, 3, 5, 1, 3},
		InputB: 3,
		Output: []bool{true, true, true, false, true},
	},
	{
		InputA: []int{12, 1, 12},
		InputB: 10,
		Output: []bool{true, false, true},
	},
}

func Test_kidsWithCandies(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := kidsWithCandies(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_kidsWithCandies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = kidsWithCandies(testCases[0].InputA, testCases[0].InputB)
	}
}
