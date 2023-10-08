package solution1832

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	{
		Input:  "thequickbrownfoxjumpsoverthelazydog",
		Output: true,
	},
	{
		Input:  "leetcode",
		Output: false,
	},
}

func Test_checkIfPangram(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := checkIfPangram(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_checkIfPangram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = checkIfPangram(testCases[0].Input)
	}
}
