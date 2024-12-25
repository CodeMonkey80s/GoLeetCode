package solution2610

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output [][]int
}{
	{
		Input: []int{1, 3, 4, 1, 2, 3, 1},
		Output: [][]int{
			{1, 3, 4, 2}, {1, 3}, {1},
		},
	},
	{
		Input:  []int{1, 2, 3, 4},
		Output: [][]int{{4, 3, 2, 1}},
	},
}

func Test_findMatrix(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {

			m1 := make(map[int]int)
			m2 := make(map[int]int)

			output := findMatrix(tc.Input)

			seq1 := slices.All(output)
			maps.Collect(seq1)
			seq2 := slices.All(tc.Output)
			maps.Collect(seq2)

			if !reflect.DeepEqual(m1, m2) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findMatrix(testCases[0].Input)
	}
}
