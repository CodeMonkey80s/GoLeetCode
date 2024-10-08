package solution3136

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	{
		Input:  "234Adas",
		Output: true,
	},
	{
		Input:  "b3",
		Output: false,
	},
	{
		Input:  "a3$e",
		Output: false,
	},
	{
		Input:  "Ya$",
		Output: false,
	},
}

func Test_isValid(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isValid(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isValid(testCases[0].Input)
	}
}
