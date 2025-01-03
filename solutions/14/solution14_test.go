package solution14

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output string
}{
	// Mandatory Test Cases
	{
		Input:  []string{"flower", "flow", "flight"},
		Output: "fl",
	},
	{
		Input:  []string{"dog", "racecar", "car"},
		Output: "",
	},
	// Additional my custom cases
	{
		Input:  []string{"monetary", "monastery", "monster"},
		Output: "mon",
	},
	{
		Input:  []string{"aca", "cba"},
		Output: "",
	},
}

func Test_longestCommonPrefix(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := longestCommonPrefix(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_longestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = longestCommonPrefix(testCases[0].Input)
	}
}
