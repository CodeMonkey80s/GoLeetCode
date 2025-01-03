package solution9

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Num    int
	Output bool
}{
	// Mandatory Test Cases
	{
		Num:    121,
		Output: true,
	},
	{
		Num:    -121,
		Output: false,
	},
	{
		Num:    10,
		Output: false,
	},
	{
		Num:    1,
		Output: true,
	},
	{
		Num:    12,
		Output: false,
	},
	{
		Num:    99,
		Output: true,
	},
	// Additional my custom cases
	{
		Num:    12321,
		Output: true,
	},
	{
		Num:    678876,
		Output: true,
	},
	{
		Num:    987656789,
		Output: true,
	},
}

func Test_isPalindrome(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Num: %v, Output: %v\n", tc.Num, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isPalindrome(tc.Num)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPalindrome(testCases[0].Num)
	}
}
