package solution1507

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
		Input:  "20th Oct 2052",
		Output: "2052-10-20",
	},
	{
		Input:  "6th Jun 1933",
		Output: "1933-06-06",
	},
	{
		Input:  "26th May 1960",
		Output: "1960-05-26",
	},
	// Additional my custom cases
}

func Test_reformatDate(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reformatDate(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_reformatDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reformatDate(testCases[0].Input)
	}
}
