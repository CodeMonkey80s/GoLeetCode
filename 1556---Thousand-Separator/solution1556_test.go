package solution1556

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output string
}{
	// Mandatory Test Cases
	{
		Input:  987,
		Output: "987",
	},
	{
		Input:  1_234,
		Output: "1.234",
	},
	{
		Input:  124_567,
		Output: "124.567",
	},
	// Additional my custom cases
	{
		Input:  123_456_789,
		Output: "123.456.789",
	},
}

func Test_thousandSeparator(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := thousandSeparator(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func Benchmark_thousandSeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = thousandSeparator(testCases[0].Input)
	}
}
