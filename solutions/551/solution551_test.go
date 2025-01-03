package solution551

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output bool
}{
	// Mandatory Test Cases
	{
		Input:  "PPALLP",
		Output: true,
	},
	{
		Input:  "PPALLL",
		Output: false,
	},
	// Additional my custom cases
}

func Test_checkRecords(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := checkRecord(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_checkRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = checkRecord(testCases[0].Input)
	}
}
