package solution2798

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
		InputA: []int{0, 1, 2, 3, 4},
		InputB: 2,
		Output: 3,
	},
	{
		InputA: []int{5, 1, 4, 2, 2},
		InputB: 6,
		Output: 0,
	},
}

func Test_numberOfEmployees(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numberOfEmployeesWhoMetTarget(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_numberOfEmployees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numberOfEmployeesWhoMetTarget(testCases[0].InputA, testCases[0].InputB)
	}
}
