package solution1678

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	// Mandatory Test Cases
	{
		Input:  "G()(al)",
		Output: "Goal",
	},
	{
		Input:  "G()()()()(al)",
		Output: "Gooooal",
	},
	{
		Input:  "(al)G(al)()()G",
		Output: "alGalooG",
	},
	// Additional my custom cases
}

func Test_interpret(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := interpret(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_interpret(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = interpret(testCases[0].Input)
	}
}
