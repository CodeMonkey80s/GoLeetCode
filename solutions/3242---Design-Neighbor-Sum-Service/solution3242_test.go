package solution3242

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB [][]int
	InputC []int
	Output []int
}{
	{
		InputA: []string{"adjacentSum", "adjacentSum", "diagonalSum", "diagonalSum"},
		InputB: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		InputC: []int{1, 4, 4, 8},
		Output: []int{6, 16, 16, 4},
	},
	{
		InputA: []string{"adjacentSum", "diagonalSum"},
		InputB: [][]int{{1, 2, 0, 3}, {4, 7, 15, 6}, {8, 9, 10, 11}, {12, 13, 14, 5}},
		InputC: []int{15, 9},
		Output: []int{23, 45},
	},
}

func Test_getFinalState(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			var output []int
			obj := Constructor(tc.InputB)
			for i, cmd := range tc.InputA {
				switch {
				case cmd == "adjacentSum":
					v := obj.AdjacentSum(tc.InputC[i])
					output = append(output, v)
				case cmd == "diagonalSum":
					v := obj.DiagonalSum(tc.InputC[i])
					output = append(output, v)
				}
			}
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkGetFinalState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := Constructor(testCases[0].InputB)
		_ = obj.AdjacentSum(testCases[0].InputC[0])
		_ = obj.DiagonalSum(testCases[0].InputC[1])
	}
}
