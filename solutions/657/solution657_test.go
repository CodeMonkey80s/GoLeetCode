package solution657

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	// Mandatory Test Cases
	{
		Input:  "UD",
		Output: true,
	},
	{
		Input:  "LL",
		Output: false,
	},
	// Additional my custom cases
}

func Test_judgeCircle(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := judgeCircle_map(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_judgeCircle_switch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = judgeCircle_switch(testCases[0].Input)
	}
}

func Benchmark_judgeCircle_map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = judgeCircle_map(testCases[0].Input)
	}
}
