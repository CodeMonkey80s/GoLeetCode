package solution2446

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB []string
	Output bool
}{
	{
		InputA: []string{"01:15", "02:00"},
		InputB: []string{"02:00", "03:00"},
		Output: true,
	},
	{
		InputA: []string{"01:00", "02:00"},
		InputB: []string{"01:20", "03:00"},
		Output: true,
	},
	{
		InputA: []string{"10:00", "11:00"},
		InputB: []string{"14:00", "15:00"},
		Output: false,
	},
	{
		InputA: []string{"14:13", "22:08"},
		InputB: []string{"02:40", "08:08"},
		Output: false,
	},
	{
		InputA: []string{"01:49", "20:39"},
		InputB: []string{"03:43", "10:26"},
		Output: true,
	},
	{
		InputA: []string{"05:10", "15:05"},
		InputB: []string{"14:59", "19:17"},
		Output: true,
	},
}

func Test_haveConflict(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := haveConflict(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_haveConflict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = haveConflict(testCases[0].InputA, testCases[0].InputB)
	}
}
