package solution3289

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{
	// Mandatory Test Cases
	{
		Input:  []int{0, 1, 1, 0},
		Output: []int{0, 1},
	},
	{
		Input:  []int{0, 3, 2, 1, 3, 2},
		Output: []int{2, 3},
	},
	{
		Input:  []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2},
		Output: []int{4, 5},
	},
	// Additional my custom cases
}

func Test_getSneakyNumbers(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getSneakyNumbersV2(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_getSneakyNumbersV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getSneakyNumbersV2(testCases[0].Input)
	}
}

func Benchmark_getSneakyNumbersV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getSneakyNumbersV1(testCases[0].Input)
	}
}
