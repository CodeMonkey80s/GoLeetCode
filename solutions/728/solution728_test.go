package solution728

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA int
	InputB int
	Output []int
}{
	{
		InputA: 1,
		InputB: 22,
		Output: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
	},
	{
		InputA: 47,
		InputB: 85,
		Output: []int{48, 55, 66, 77},
	},
}

func Test_selfDividingNumbers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := selfDividingNumbers(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_selfDividingNumbers(b *testing.B) {
	for b.Loop() {
		_ = selfDividingNumbers(testCases[0].InputA, testCases[0].InputB)
	}
}
