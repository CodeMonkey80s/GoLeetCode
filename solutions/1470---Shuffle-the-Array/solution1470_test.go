package solution1470

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
		InputA: []int{2, 5, 1, 3, 4, 7},
		InputB: 3,
		Output: []int{2, 3, 5, 4, 1, 7},
	},
	{
		InputA: []int{1, 2, 3, 4, 4, 3, 2, 1},
		InputB: 4,
		Output: []int{1, 4, 2, 3, 3, 2, 4, 1},
	},
}

func Test_shuffl(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := shuffle(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_shuffle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shuffle(testCases[0].InputA, testCases[0].InputB)
	}
}
