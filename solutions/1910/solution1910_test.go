package solution1910

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output string
}{
	{
		InputA: "daabcbaabcbc",
		InputB: "abc",
		Output: "dab",
	},
	{
		InputA: "axxxxyyyyb",
		InputB: "xy",
		Output: "ab",
	},
}

func Test_removeOccurrences(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := removeOccurrencesV2(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_removeOccurrencesV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeOccurrencesV1(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_removeOccurrencesV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeOccurrencesV2(testCases[0].InputA, testCases[0].InputB)
	}
}
