package solution961

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 2, 3, 3},
		Output: 3,
	},
	{
		Input:  []int{2, 1, 2, 5, 3, 2},
		Output: 2,
	},
	{
		Input:  []int{5, 1, 5, 2, 5, 3, 5, 4},
		Output: 5,
	},
}

func Test_repeatedNTimes(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := repeatedNTimesV2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_repeatedNTimesV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = repeatedNTimesV2(testCases[0].Input)
	}
}

func Benchmark_repeatedNTimesV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = repeatedNTimesV1(testCases[0].Input)
	}
}
