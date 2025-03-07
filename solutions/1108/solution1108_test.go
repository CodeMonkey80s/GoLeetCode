package solution1108

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	// Mandatory Test Cases
	{
		Input:  "1.1.1.1",
		Output: "1[.]1[.]1[.]1",
	},
	{
		Input:  "255.100.50.0",
		Output: "255[.]100[.]50[.]0",
	},
	// Additional my custom cases
}

func Test_defang(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := defangIPaddrV2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_defangV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = defangIPaddrV2(testCases[0].Input)
	}
}

func Benchmark_defang_stdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = defangIPaddrV1(testCases[0].Input)
	}
}
