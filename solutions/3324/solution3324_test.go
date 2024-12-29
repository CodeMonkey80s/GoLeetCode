package solution3324

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output []string
}{
	{
		Input:  "abc",
		Output: []string{"a", "aa", "ab", "aba", "abb", "abc"},
	},
	{
		Input:  "he",
		Output: []string{"a", "b", "c", "d", "e", "f", "g", "h", "ha", "hb", "hc", "hd", "he"},
	},
}

func Test_stringSequence(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := stringSequence(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_stringSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stringSequence(testCases[1].Input)
	}
}
