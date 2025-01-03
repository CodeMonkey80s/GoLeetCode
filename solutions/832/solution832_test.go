package solution832

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output [][]int
}{
	{
		Input: [][]int{
			{1, 1, 0},
			{1, 0, 1},
			{0, 0, 0},
		},
		Output: [][]int{
			{1, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
	},
	{
		Input: [][]int{
			{1, 1, 0, 0},
			{1, 0, 0, 1},
			{0, 1, 1, 1},
			{1, 0, 1, 0},
		},
		Output: [][]int{
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 1},
			{1, 0, 1, 0},
		},
	},
}

func Test_flipAndInvertImage(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := flipAndInvertImage(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_flipAndInvertImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = flipAndInvertImage(testCases[0].Input)
	}
}
