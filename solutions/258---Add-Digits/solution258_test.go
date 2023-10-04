package solution258

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Num    int
	Output int
}{
	// Mandatory Test Cases
	{
		Num:    38,
		Output: 2,
	},
	{
		Num:    0,
		Output: 0,
	},
	// Additional my custom cases
	{
		Num:    123,
		Output: 6,
	},
	{
		Num:    48671,
		Output: 8,
	},
}

func Test_addDigits(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Num: %v, Output: %v\n", tc.Num, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := addDigits(tc.Num)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_addDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addDigits(testCases[0].Num)
	}
}
