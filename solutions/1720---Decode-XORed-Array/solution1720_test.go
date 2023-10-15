package solution1720

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
		InputA: []int{1, 2, 3},
		InputB: 1,
		Output: []int{1, 0, 2, 1},
	},
	{
		InputA: []int{6, 2, 7, 3},
		InputB: 4,
		Output: []int{4, 2, 0, 7, 4},
	},
}

func Test_decode(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := decode(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = decode(testCases[0].InputA, testCases[0].InputB)
	}
}
