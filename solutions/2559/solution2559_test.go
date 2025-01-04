package solution2559

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB [][]int
	Output []int
}{
	{
		InputA: []string{"aba", "bcb", "ece", "aa", "e"},
		InputB: [][]int{{0, 2}, {1, 4}, {1, 1}},
		Output: []int{2, 3, 0},
	},
	{
		InputA: []string{"a", "e", "i"},
		InputB: [][]int{{0, 2}, {0, 1}, {2, 2}},
		Output: []int{3, 2, 1},
	},
}

func Test_vowelString(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := vowelString(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_vowelString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = vowelString(testCases[0].InputA, testCases[0].InputB)
	}
}
