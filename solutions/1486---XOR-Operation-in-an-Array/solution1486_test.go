package solution1486

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
		InputA: 5,
		InputB: 0,
		Output: 8,
	},
	{
		InputA: 4,
		InputB: 3,
		Output: 8,
	},
}

func Test_xorOperation(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := xorOperation(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_xorOperation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = xorOperation(testCases[0].InputA, testCases[0].InputB)
	}
}
