package solution349

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA  []int
	InputB  []int
	OutputA []int
	OutputB []int
}{
	// Mandatory Test Cases
	{
		InputA:  []int{1, 2, 2, 1},
		InputB:  []int{2, 2},
		OutputA: []int{2},
		OutputB: []int{},
	},
	{
		InputA:  []int{4, 9, 5},
		InputB:  []int{9, 4, 9, 8, 3},
		OutputA: []int{4, 9},
		OutputB: []int{9, 4},
	},
}

func Test_intersection(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.OutputA)
		t.Run(label, func(t *testing.T) {
			output := intersection(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.OutputA) && !reflect.DeepEqual(output, tc.OutputB) {
				t.Errorf("Expected output to be %v or %v but we got %v", tc.OutputA, tc.OutputB, output)
			}
		})
	}
}

func Benchmark_intersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = intersection(testCases[0].InputA, testCases[0].InputB)
	}
}
