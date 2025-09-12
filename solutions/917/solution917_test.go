package solution917

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "ab-cd",
		Output: "dc-ba",
	},
	{
		Input:  "a-bC-dEf-ghIj",
		Output: "j-Ih-gfE-dCba",
	},
	{
		Input:  "Test1ng-Leet=code-Q!",
		Output: "Qedo1ct-eeLg=ntse-T!",
	},
}

func Test_reverseOnlyLetters(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reverseOnlyLetters(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_reverseOnlyLetters(b *testing.B) {
	for b.Loop() {
		_ = reverseOnlyLetters(testCases[0].Input)
	}
}
