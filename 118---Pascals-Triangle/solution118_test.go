package solution118

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  int
	Output [][]int
}{
	// Mandatory Test Cases
	{
		Input: 5,
		Output: [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		},
	},
	// Additional my custom cases
	{
		Input: 1,
		Output: [][]int{
			{1},
		},
	},
	{
		Input: 2,
		Output: [][]int{
			{1},
			{1, 1},
		},
	},
	{
		Input: 3,
		Output: [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
		},
	},
}

func Test_generate(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := generate(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_generate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate(testCases[0].Input)
	}
}
