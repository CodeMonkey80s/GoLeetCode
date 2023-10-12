package solution2194

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output []string
}{
	{
		Input:  "K1:L2",
		Output: []string{"K1", "K2", "L1", "L2"},
	},
	{
		Input:  "A1:F1",
		Output: []string{"A1", "B1", "C1", "D1", "E1", "F1"},
	},
	{
		Input:  "U7:V9",
		Output: []string{"U7", "U8", "U9", "V7", "V8", "V9"},
	},
}

func Test_cellsInRange(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := cellsInRange(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_cellsInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cellsInRange(testCases[0].Input)
	}
}

func Benchmark_cellsInRange_string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cellsInRange_string(testCases[0].Input)
	}
}
