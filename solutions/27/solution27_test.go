package solution27

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output int
}{
	{
		InputA: []int{3, 2, 2, 3},
		InputB: 3,
		Output: 2,
	},
	{
		InputA: []int{0, 1, 2, 2, 3, 0, 4, 2},
		InputB: 2,
		Output: 5,
	},
	{
		InputA: []int{4, 4, 0, 1, 0, 2},
		InputB: 0,
		Output: 4,
	},
	{
		InputA: []int{},
		InputB: 0,
		Output: 0,
	},
}

func Test_removeElement(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Remove: %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := removeElement(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_removeElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeElement(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_removeElementV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeElementV1(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_removeElement_stdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeElement_stdlib(testCases[0].InputA, testCases[0].InputB)
	}
}
