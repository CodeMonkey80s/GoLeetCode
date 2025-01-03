package solution1360

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output int
}{
	// Mandatory Test Cases
	{
		InputA: "2019-06-29",
		InputB: "2019-06-30",
		Output: 1,
	},
	{
		InputA: "2020-01-15",
		InputB: "2019-12-31",
		Output: 15,
	},
	// Additional my custom cases
}

func Test_daysBetween(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := daysBetweenDates(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_daysBetween(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = daysBetweenDates(testCases[0].InputA, testCases[0].InputB)
	}
}
