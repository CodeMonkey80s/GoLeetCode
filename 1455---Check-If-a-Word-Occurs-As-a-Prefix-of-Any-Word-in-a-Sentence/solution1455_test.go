package solution1455

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output int
}{
	// Mandatory Test Cases
	{
		InputA: "i love eating burger",
		InputB: "burg",
		Output: 4,
	},
	{
		InputA: "this problem is an easy problem",
		InputB: "pro",
		Output: 2,
	},
	{
		InputA: "i am tired",
		InputB: "you",
		Output: -1,
	},
	// Additional my custom cases
	{
		InputA: "hellohello hellohellohello",
		InputB: "ell",
		Output: -1,
	},
}

func Test_isPrefixOfWord(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isPrefixOfWord(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isPrefixOfWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPrefixOfWord(testCases[0].InputA, testCases[0].InputB)
	}
}
