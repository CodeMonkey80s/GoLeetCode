package solution125

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	// Mandatory Test Cases
	{
		Input:  "A man, a plan, a canal: Panama",
		Output: true,
	},
	{
		Input:  "race a car",
		Output: false,
	},
	{
		Input:  "",
		Output: true,
	},
	// Additional my custom cases
	{
		Input:  "0P",
		Output: false,
	},
	{
		Input:  "a",
		Output: true,
	},
}

func Test_isPalindrome(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isPalindrome(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPalindrome(testCases[0].Input)
	}
}
