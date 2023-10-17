package solution1089

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{
	{
		Input:  []int{1, 0, 2, 3, 0, 4, 5, 0},
		Output: []int{1, 0, 0, 2, 3, 0, 0, 4},
	},
	{
		Input:  []int{1, 2, 3},
		Output: []int{1, 2, 3},
	},
}

func Test_duplicateZeros(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			duplicateZeros(tc.Input)
			if !reflect.DeepEqual(tc.Input, tc.Output) {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, tc.Input)
			}
		})
	}
}

func Benchmark_duplicateZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		duplicateZeros(testCases[0].Input)
	}
}
