package solution2006

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output int
}{
	{
		InputA: []int{1, 2, 2, 1},
		InputB: 1,
		Output: 4,
	},
	{
		InputA: []int{1, 3},
		InputB: 3,
		Output: 0,
	},
}

func Test_countKDifference(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countKDifference(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_count(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countKDifference(testCases[0].InputA, testCases[0].InputB)
	}
}
