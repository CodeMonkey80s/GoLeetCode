package solution338

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  int
	Output []int
}{
	// Mandatory Test Cases
	{
		Input:  2,
		Output: []int{0, 1, 1},
	},
	{
		Input:  5,
		Output: []int{0, 1, 1, 2, 1, 2},
	},
	// Additional my custom cases
}

func Test_countBits(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countBitsV2(tc.Input)
			if !reflect.DeepEqual(tc.Output, output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countBitsV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countBitsV2(testCases[0].Input)
	}
}

func Benchmark_countBitsV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countBitsV1(testCases[0].Input)
	}
}
