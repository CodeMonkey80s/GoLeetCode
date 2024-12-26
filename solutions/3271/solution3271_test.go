package solution3271

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB int
	Output string
}{
	{
		InputA: "abcd",
		InputB: 2,
		Output: "bf",
	},
	{
		InputA: "mxz",
		InputB: 3,
		Output: "i",
	},
}

func Test_findMatrix(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := stringHashV1(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_stringHashV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stringHashV1(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_stringHashV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stringHashV2(testCases[0].InputA, testCases[0].InputB)
	}
}
