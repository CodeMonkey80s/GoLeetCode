package solution3354

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{

	{
		Input:  []int{1, 0, 2, 0, 3},
		Output: 2,
	},
	{
		Input:  []int{2, 3, 4, 0, 4, 1, 0},
		Output: 0,
	},
	{
		Input:  []int{0},
		Output: 2,
	},
	{
		Input:  []int{16, 13, 10, 0, 0, 0, 10, 6, 7, 8, 7},
		Output: 3,
	},
}

func Test_countValidSelections(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countValidSelectionsV1(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countValidSelectionsV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countValidSelectionsV1(testCases[0].Input)
	}
}

func Benchmark_countValidSelectionsV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countValidSelectionsV2(testCases[0].Input)
	}
}

func Benchmark_countValidSelectionsV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countValidSelectionsV3(testCases[0].Input)
	}
}
