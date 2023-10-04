package solution405

import (
	"fmt"
	"strconv"
	"testing"
)

var testCases = []struct {
	Input  int
	Output string
}{
	{
		Input:  26,
		Output: "1a",
	},
	{
		Input:  13485,
		Output: "34ad",
	},
	{
		Input:  0,
		Output: "0",
	},
	{
		Input:  -1,
		Output: "ffffffff",
	},
}

func Test_toHex(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := toHex(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_toHex_my(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toHex(testCases[0].Input)
	}
}

func Benchmark_toHex_std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(int64(testCases[0].Input), 16)
	}
}
