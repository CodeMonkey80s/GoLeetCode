package solution2215

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	Output [][]int
}{
	{
		InputA: []int{1, 2, 3},
		InputB: []int{2, 4, 6},
		Output: [][]int{{1, 3}, {4, 6}},
	},
	{
		InputA: []int{1, 2, 3, 3},
		InputB: []int{1, 1, 2, 2},
		Output: [][]int{{3}, {}},
	},
	{
		InputA: []int{26, 48, -78, -25, 42, -8, 94, -68, 26},
		InputB: []int{61, -17},
		Output: [][]int{{26, 48, -78, -25, 42, -8, 94, -68}, {61, -17}},
	},
}

func Test_findDifference(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findDifference(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findDifference(testCases[0].InputA, testCases[0].InputB)
	}
}
