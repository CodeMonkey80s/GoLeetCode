package solution2788

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB byte
	Output []string
}{
	{
		InputA: []string{"one.two.three", "four.five", "six"},
		InputB: '.',
		Output: []string{"one", "two", "three", "four", "five", "six"},
	},
	{
		InputA: []string{"$easy$", "$problem$"},
		InputB: '$',
		Output: []string{"easy", "problem"},
	},
	{
		InputA: []string{"|||"},
		InputB: '|',
		Output: []string{},
	},
}

func Test_splitWords(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := splitWordsBySeparator(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_splitWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = splitWordsBySeparator(testCases[0].InputA, testCases[0].InputB)
	}
}
