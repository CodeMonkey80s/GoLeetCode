package solution228

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []string
}{
	{
		Input:  []int{0, 1, 2, 4, 5, 7},
		Output: []string{"0->2", "4->5", "7"},
	},
	{
		Input:  []int{0, 2, 3, 4, 6, 8, 9},
		Output: []string{"0", "2->4", "6", "8->9"},
	},
}

func Test_summaryRanges(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := summaryRanges(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_summaryRanges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = summaryRanges(testCases[0].Input)
	}
}
