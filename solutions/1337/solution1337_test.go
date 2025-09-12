package solution1337

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA [][]int
	InputB int
	Output []int
}{
	{
		InputA: [][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		},
		InputB: 3,
		Output: []int{2, 0, 3},
	},
	{
		InputA: [][]int{
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		},
		InputB: 2,
		Output: []int{0, 2},
	},
}

func Test_KWeakestRows(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := kWeakestRows(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_KWeakestRows(b *testing.B) {
	for b.Loop() {
		_ = kWeakestRows(testCases[0].InputA, testCases[0].InputB)
	}
}
