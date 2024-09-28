package solution3264

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	InputC int
	Output []int
}{
	{
		InputA: []int{2, 1, 3, 5, 6},
		InputB: 5,
		InputC: 2,
		Output: []int{8, 4, 6, 5, 6},
	},
	{
		InputA: []int{1, 2},
		InputB: 3,
		InputC: 4,
		Output: []int{16, 8},
	},
	{
		InputA: []int{7, 71, 15},
		InputB: 10,
		InputC: 3,
		Output: []int{567, 639, 1215},
	},
}

func Test_getFinalState(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getFinalState(tc.InputA, tc.InputB, tc.InputC)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkGetFinalState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getFinalState(testCases[0].InputA, testCases[0].InputB, testCases[0].InputC)
	}
}
