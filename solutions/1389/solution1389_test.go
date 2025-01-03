package solution1389

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
		InputA: []int{0, 1, 2, 3, 4},
		InputB: []int{0, 1, 2, 2, 1},
		Output: []int{0, 4, 1, 3, 2},
	},
	{
		InputA: []int{1, 2, 3, 4, 0},
		InputB: []int{0, 1, 2, 3, 0},
		Output: []int{0, 1, 2, 3, 4},
	},
}

func Test_createTargetArray(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := createTargetArray(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_createTargetArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = createTargetArray(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_createTargetArray_std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = createTargetArray_std(testCases[0].InputA, testCases[0].InputB)
	}
}
