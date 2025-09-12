package solution1544

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
		Input:  "leEeetcode",
		Output: "leetcode",
	},
	{
		Input:  "abBAcC",
		Output: "",
	},
	{
		Input:  "s",
		Output: "s",
	},
}

func Test_makeGood(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := makeGood(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_makeGood(b *testing.B) {
	for b.Loop() {
		_ = makeGood(testCases[0].Input)
	}
}
