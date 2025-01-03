package solution2169

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA int
	InputB int
	Output int
}{
	{
		InputA: 2,
		InputB: 3,
		Output: 3,
	},
	{
		InputA: 10,
		InputB: 10,
		Output: 1,
	},
}

func Test_countOperations(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countOperations(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_CountOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countOperations(testCases[0].InputA, testCases[0].InputB)
	}
}
