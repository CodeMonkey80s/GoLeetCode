package solution13

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  "III",
		Output: 3,
	},
	{
		Input:  "LVIII",
		Output: 58,
	},
	{
		Input:  "MCMXCIV",
		Output: 1994,
	},
	// Additional my custom cases
	{
		Input:  "MC",
		Output: 1100,
	},
	{
		Input:  "MIV",
		Output: 1004,
	},
	{
		Input:  "CXC",
		Output: 190,
	},
}

func Test_romanToInt(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := romanToInt(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_romanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = romanToInt(testCases[0].Input)
	}
}
