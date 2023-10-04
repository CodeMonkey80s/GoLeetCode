package solution191

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  uint32
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  11,
		Output: 3,
	},
	{
		Input:  128,
		Output: 1,
	},
	{
		Input:  4294967293,
		Output: 31,
	},
	// Additional my custom cases
}

func Test_hammingWeight_my(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := hammingWeight(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_hammingWeight_my(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hammingWeight(testCases[0].Input)
	}
}

func Benchmark_hammingWeight_stdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hammingWeight_stdlib(testCases[0].Input)
	}
}
