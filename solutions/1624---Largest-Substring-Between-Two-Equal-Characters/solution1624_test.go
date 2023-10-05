package solution1624

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "aa",
		Output: 0,
	},
	{
		Input:  "abca",
		Output: 2,
	},
	{
		Input:  "cbzxy",
		Output: -1,
	},
}

func Test_maxLength(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maxLengthBetweenEqualCharacters(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maxLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxLengthBetweenEqualCharacters(testCases[0].Input)
	}
}
