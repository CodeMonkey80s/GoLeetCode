package solution709

import (
	"fmt"
	"strings"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	// Mandatory Test Cases
	{
		Input:  "Hello",
		Output: "hello",
	},
	{
		Input:  "here",
		Output: "here",
	},
	{
		Input:  "LOVELY",
		Output: "lovely",
	},
	// Additional my custom cases
}

func Test_toLowerCase(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := toLowerCase(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_toLowerCase_my(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toLowerCase(testCases[0].Input)
	}
}

func Benchmark_toLowerCase_std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.ToLower(testCases[0].Input)
	}
}
