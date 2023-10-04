package solution67

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output string
}{
	// Mandatory Test Cases
	{
		InputA: "11",
		InputB: "1",
		Output: "100",
	},
	{
		InputA: "1010",
		InputB: "1011",
		Output: "10101",
	},
	// Additional my custom cases
	{
		InputA: "1111",
		InputB: "1111",
		Output: "11110",
	},
	{
		InputA: "0",
		InputB: "0",
		Output: "0",
	},
	{
		InputA: "0",
		InputB: "1",
		Output: "1",
	},
	{
		InputA: "100",
		InputB: "110010",
		Output: "110110",
	},
}

func Test_addBinary(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := addBinary_my(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_addBinary_stdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addBinary_stdlib(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_addBinary_my(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addBinary_my(testCases[0].InputA, testCases[0].InputB)
	}
}
