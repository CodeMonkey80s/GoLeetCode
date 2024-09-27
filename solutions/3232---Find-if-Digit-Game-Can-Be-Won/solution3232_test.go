package solution3232

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output bool
}{
	{
		Input:  []int{1, 2, 3, 4, 10},
		Output: false,
	},
	{
		Input:  []int{1, 2, 3, 4, 5, 14},
		Output: true,
	},
	{
		Input:  []int{5, 5, 5, 25},
		Output: true,
	},
}

func Test_canAliceWin(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := canAliceWin(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_canAliceWing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = canAliceWin(testCases[0].Input)
	}
}
