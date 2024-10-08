package solution476

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output int
}{
	{
		Input:  5,
		Output: 2,
	},
	{
		Input:  1,
		Output: 0,
	},
	{
		Input:  2,
		Output: 1,
	},
}

func Test_findComplement(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findComplementV3(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findComplementV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findComplementV3(testCases[0].Input)
	}
}

func Benchmark_findComplementV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findComplementV2(testCases[0].Input)
	}
}

func Benchmark_findComplementV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findComplementV1(testCases[0].Input)
	}
}
