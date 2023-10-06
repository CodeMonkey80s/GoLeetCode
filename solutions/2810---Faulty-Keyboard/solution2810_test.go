package solution2810

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "string",
		Output: "rtsng",
	},
	{
		Input:  "poiinter",
		Output: "ponter",
	},
	{
		Input:  "fii",
		Output: "f",
	},
	{
		Input:  "qskyviiiii",
		Output: "vyksq",
	},
	{
		Input:  "viwif",
		Output: "wvf",
	},
}

func Test_finalString(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := finalString(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_finalString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = finalString(testCases[0].Input)
	}
}
