package solution168

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Num    int
	Output string
}{
	// Mandatory Test Cases
	{
		Num:    1,
		Output: "A",
	},
	{
		Num:    28,
		Output: "AB",
	},
	{
		Num:    701,
		Output: "ZY",
	},
	// Additional my custom cases
}

func Test_convertToTitle(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Num: %v, Output: %v\n", tc.Num, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := convertToTitle(tc.Num)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_convertTotTile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = convertToTitle(testCases[0].Num)
	}
}
