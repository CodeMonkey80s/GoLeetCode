package solution2864

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "010",
		Output: "001",
	},
	{
		Input:  "0101",
		Output: "1001",
	},
}

func Test_maximumOddBinary(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maximumOddBinaryNumber(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maximumOddBinaryNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maximumOddBinaryNumber(testCases[0].Input)
	}
}
