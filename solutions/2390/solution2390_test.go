package solution2390

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
		Input:  "leet**cod*e",
		Output: "lecoe",
	},
	{
		Input:  "erase*****",
		Output: "",
	},
}

func Test_removeStars(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := removeStars(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_removeStars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeStars(testCases[0].Input)
	}
}
