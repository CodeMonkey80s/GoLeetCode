package solution1184

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	InputC int
	Output int
}{
	{
		InputA: []int{1, 2, 3, 4},
		InputB: 0,
		InputC: 1,
		Output: 1,
	},
	{
		InputA: []int{1, 2, 3, 4},
		InputB: 0,
		InputC: 2,
		Output: 3,
	},
	{
		InputA: []int{1, 2, 3, 4},
		InputB: 0,
		InputC: 3,
		Output: 4,
	},
	{
		InputA: []int{7, 10, 1, 12, 11, 14, 5, 0},
		InputB: 7,
		InputC: 2,
		Output: 17,
	},
}

func Test_distance(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := distanceBetweenBusStopsV2(tc.InputA, tc.InputB, tc.InputC)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_distance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = distanceBetweenBusStopsV2(testCases[0].InputA, testCases[0].InputB, testCases[0].InputC)
	}
}
