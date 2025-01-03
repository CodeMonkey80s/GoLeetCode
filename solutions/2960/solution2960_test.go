package solution2960

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 1, 2, 1, 3},
		Output: 3,
	},
	{
		Input:  []int{0, 1, 2},
		Output: 2,
	},
}

func Test_countTestedDevices(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countTestedDevicesV2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_CountTestedDevicesV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countTestedDevicesV2(testCases[0].Input)
	}
}

func Benchmark_CountTestedDevicesV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countTestedDevicesV1(testCases[0].Input)
	}
}
