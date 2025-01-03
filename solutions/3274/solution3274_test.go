package solution3274

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output bool
}{

	{
		InputA: "a1",
		InputB: "c3",
		Output: true,
	},
	{
		InputA: "a1",
		InputB: "h3",
		Output: false,
	},
}

func Test_checkTwoChessboards(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := checkTwoChessboards(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_checkTwoChessboards(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = checkTwoChessboards(testCases[0].InputA, testCases[0].InputB)
	}
}
