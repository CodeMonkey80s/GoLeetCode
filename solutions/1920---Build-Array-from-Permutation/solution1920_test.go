package solution1920

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{
	{
		Input:  []int{0, 2, 1, 5, 3, 4},
		Output: []int{0, 1, 2, 4, 5, 3},
	},
	{
		Input:  []int{5, 0, 1, 2, 3, 4},
		Output: []int{4, 5, 0, 1, 2, 3},
	},
}

func Test_buildArray(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := buildArray(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_buildArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = buildArray(testCases[0].Input)
	}
}
