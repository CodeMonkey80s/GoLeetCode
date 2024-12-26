package solution2120

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA int
	InputB []int
	InputC string
	Output []int
}{
	{
		InputA: 3,
		InputB: []int{0, 1},
		InputC: "RRDDLU",
		Output: []int{1, 5, 4, 3, 1, 0},
	},
	{
		InputA: 2,
		InputB: []int{1, 1},
		InputC: "LURD",
		Output: []int{4, 1, 0, 0},
	},
	{
		InputA: 1,
		InputB: []int{0, 0},
		InputC: "LURD",
		Output: []int{0, 0, 0, 0},
	},
}

func Test_executeInstructions(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := executeInstructions(tc.InputA, tc.InputB, tc.InputC)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_executeInstructions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = executeInstructions(testCases[0].InputA, testCases[1].InputB, testCases[2].InputC)
	}
}
