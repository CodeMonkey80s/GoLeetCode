package solution119

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  int
	Output []int
}{
	// Mandatory Test Cases
	{
		Input:  3,
		Output: []int{1, 3, 3, 1},
	},
	{
		Input:  0,
		Output: []int{1},
	},
	{
		Input:  1,
		Output: []int{1, 1},
	},
	// Additional my custom cases
}

func Test_getRow(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getRow(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_getRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getRow(testCases[0].Input)
	}
}
