package solution415

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output string
}{
	{
		InputA: "11",
		InputB: "123",
		Output: "134",
	},
	{
		InputA: "456",
		InputB: "77",
		Output: "533",
	},
	{
		InputA: "9",
		InputB: "99",
		Output: "108",
	},
	{
		InputA: "1",
		InputB: "1",
		Output: "2",
	},
	{
		InputA: "9133",
		InputB: "0",
		Output: "9133",
	},
	{
		InputA: "82978",
		InputB: "75",
		Output: "83053",
	},
	{
		InputA: "6",
		InputB: "82503",
		Output: "82509",
	},
	{
		InputA: "0",
		InputB: "0",
		Output: "0",
	},
}

func Test_addStrings(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := addStrings_std(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_addStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addStrings(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_addStrings_std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addStrings_std(testCases[0].InputA, testCases[0].InputB)
	}
}
