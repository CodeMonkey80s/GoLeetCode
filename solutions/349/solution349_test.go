package solution349

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	Output []int
}{
	{
		InputA: []int{1, 2, 2, 1},
		InputB: []int{2, 2},
		Output: []int{2},
	},
	{
		InputA: []int{4, 9, 5},
		InputB: []int{9, 4, 9, 8, 4},
		Output: []int{4, 9},
	},
	{
		InputA: []int{8, 0, 3},
		InputB: []int{0, 0},
		Output: []int{0},
	},
	{
		InputA: []int{4, 7, 9, 7, 6, 7},
		InputB: []int{5, 0, 0, 6, 1, 6, 2, 2, 4},
		Output: []int{4, 6},
	},
}

func Test_intersection(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := intersectionV2(tc.InputA, tc.InputB)
			slices.Sort(output)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_intersectionV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = intersectionV2(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_intersectionV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = intersectionV1(testCases[0].InputA, testCases[0].InputB)
	}
}
