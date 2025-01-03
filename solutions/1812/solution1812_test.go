package solution1812

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	{
		Input:  "a1",
		Output: false,
	},
	{
		Input:  "h3",
		Output: true,
	},
	{
		Input:  "c7",
		Output: false,
	},
	{
		Input:  "h8",
		Output: false,
	},
	{
		Input:  "a8",
		Output: true,
	},
}

func Test_squareIsWhite(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := squareIsWhiteV2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_squareIsWhiteV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = squareIsWhiteV2(testCases[0].Input)
	}
}

func Benchmark_squareIsWhiteV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = squareIsWhiteV1(testCases[0].Input)
	}
}
