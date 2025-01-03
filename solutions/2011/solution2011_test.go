package solution2011

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input:  []string{"--X", "X++", "X++"},
		Output: 1,
	},
	{
		Input:  []string{"++X", "++X", "X++"},
		Output: 3,
	},
	{
		Input:  []string{"X++", "++X", "--X", "X--"},
		Output: 0,
	},
}

func Test_finalValue(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := finalValueAfterOperations(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_finalValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = finalValueAfterOperations(testCases[2].Input)
	}
}

func Benchmark_finalValue_if(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = finalValueAfterOperations_if(testCases[2].Input)
	}
}
