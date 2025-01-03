package solution884

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	Output []string
}{
	// Mandatory Test Cases
	{
		InputA: "this apple is sweet",
		InputB: "this apple is sour",
		Output: []string{"sweet", "sour"},
	},
	{
		InputA: "apple apple",
		InputB: "banana",
		Output: []string{"banana"},
	},
	// Additional my custom cases
	{
		InputA: "s z z z s",
		InputB: "s z ejt",
		Output: []string{"ejt"},
	},
}

func Test_uncommonFromSentences(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := uncommonFromSentences(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_uncommonFromSentences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uncommonFromSentences(testCases[0].InputA, testCases[0].InputB)
	}
}
