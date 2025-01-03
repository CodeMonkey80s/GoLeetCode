package solution2540

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB []int
	Output int
}{
	{
		InputA: []int{1, 2, 3},
		InputB: []int{2, 4},
		Output: 2,
	},
	{
		InputA: []int{1, 2, 3, 6},
		InputB: []int{2, 3, 4, 5},
		Output: 2,
	},
	{
		InputA: []int{2, 4},
		InputB: []int{1, 2},
		Output: 2,
	},
	{
		InputA: []int{1000000000, 1000000000},
		InputB: []int{1000000000},
		Output: 1000000000,
	},
}

func Test_getCommon(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getCommon(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_similarPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getCommon(testCases[0].InputA, testCases[0].InputB)
	}
}
