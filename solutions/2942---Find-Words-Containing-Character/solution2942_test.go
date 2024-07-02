package solution2942

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB byte
	Output []int
}{
	{
		InputA: []string{"leet", "code"},
		InputB: 'e',
		Output: []int{0, 1},
	},
	{
		InputA: []string{"abc", "bcd", "aaaa", "cbc"},
		InputB: 'a',
		Output: []int{0, 2},
	},
	{
		InputA: []string{"abc", "bcd", "aaaa", "cbc"},
		InputB: 'z',
		Output: []int{},
	},
}

func Test_findWordsContaining(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findWordsContaining(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findWordsContaining(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findWordsContaining(testCases[0].InputA, testCases[0].InputB)
	}
}
