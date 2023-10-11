package solution2828

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB string
	Output bool
}{
	{
		InputA: []string{"alice", "bob", "charlie"},
		InputB: "abc",
		Output: true,
	},
	{
		InputA: []string{"an", "apple"},
		InputB: "a",
		Output: false,
	},
	{
		InputA: []string{"never", "gonna", "give", "up", "on", "you"},
		InputB: "ngguoy",
		Output: true,
	},
}

func Test_isAcronym(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isAcronym(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_isAcronym(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isAcronym(testCases[0].InputA, testCases[0].InputB)
	}
}
