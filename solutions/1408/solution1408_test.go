package solution1408

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output []string
}{
	{
		Input:  []string{"mass", "as", "hero", "superhero"},
		Output: []string{"as", "hero"},
	},
	{
		Input:  []string{"leetcode", "et", "code"},
		Output: []string{"et", "code"},
	},
	{
		Input:  []string{"blue", "green", "bu"},
		Output: []string{},
	},
}

func Test_stringMatching(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := stringMatching(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_stringMatching(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stringMatching(testCases[0].Input)
	}
}
