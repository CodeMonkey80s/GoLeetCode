package solution482

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB int
	Output string
}{
	// Mandatory Test Cases
	{
		InputA: "5F3Z-2e-9-w",
		InputB: 4,
		Output: "5F3Z-2E9W",
	},
	{
		InputA: "2-5g-3-J",
		InputB: 2,
		Output: "2-5G-3J",
	},
	{
		InputA: "--a-a-a-a--",
		InputB: 2,
		Output: "AA-AA",
	},
	{
		InputA: "---",
		InputB: 2,
		Output: "",
	},
	// Additional my custom cases
}

func Test_licenseKeyFormatting(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := licenseKeyFormatting(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_licenseKeyFormatting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = licenseKeyFormatting(testCases[0].InputA, testCases[0].InputB)
	}
}
