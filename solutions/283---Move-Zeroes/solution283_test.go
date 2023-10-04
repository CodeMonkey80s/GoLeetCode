package solution283

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{
	// Mandatory Test Cases
	{
		Input:  []int{0, 1, 0, 3, 12},
		Output: []int{1, 3, 12, 0, 0},
	},
	{
		Input:  []int{0},
		Output: []int{0},
	},
	// Additional my custom cases
}

func Test_moveZeroes(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			input := make([]int, len(tc.Input))
			copy(input, tc.Input)
			moveZeroes(input)
			if !reflect.DeepEqual(input, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, input)
			}
		})
	}
}

func Benchmark_moveZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moveZeroes(testCases[0].Input)
	}
}
